// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"internal/abi"
	"internal/runtime/atomic"
	"unsafe"
)

func osyield()

//go:nosplit
func osyield_no_g() {
	osyield()
}

func osinit() {
	ncpu = int32(getnCPU())
	if physPageSize == 0 {
		physPageSize = uintptr(getPageSize())
	}
}

func setThreadCPUProfiler(hz int32) {
}

func setProcessCPUProfiler(hz int32)

//go:nosplit
func futexwakeup(addr *uint32, cnt uint32)

func futexsleep(addr *uint32, val uint32, ns int64)

func goenvs() {
	goenvs_unix()
}

// Called to initialize a new m (including the bootstrap m).
// Called on the parent thread (main thread in case of bootstrap), can allocate memory.
func mpreinit(mp *m) {
	mp.gsignal = malg(32 * 1024)
	mp.gsignal.m = mp
}

// Called to initialize a new m (including the bootstrap m).
// Called on the new thread, cannot allocate memory.
func minit() {
	getg().m.procid = gettid()
}

// Called from dropm to undo the effect of an minit.
//
//go:nosplit
func unminit() {
	getg().m.procid = 0
}

// Called from exitm, but not from drop, to undo the effect of thread-owned
// resources in minit, semacreate, or elsewhere. Do not take locks after calling this.
func mdestroy(mp *m) {}

func threadStart(f, s, a uintptr) int

func threadInit()

// May run with m.p==nil, so write barriers are not allowed.
//
//go:nowritebarrierrec
func newosproc(mp *m) {
	stk := unsafe.Pointer(mp.g0.stack.hi)
	threadStart(abi.FuncPCABI0(threadInit), uintptr(stk), uintptr(unsafe.Pointer(mp)))
}

func readRandom(r []byte) int {
	return 0
}

func signalM(mp *m, sig int) {
	kill(0, mp.procid, uint32(sig))
}

func validSIGPROF(mp *m, c *sigctxt) bool {
	return true
}

// read calls the read system call.
// It returns a non-negative number of bytes written or a negative errno value.
func read(fd int32, p unsafe.Pointer, n int32) int32

func closefd(fd int32) int32

func exit(code int32)

func usleep(usec uint32)

//go:nosplit
func usleep_no_g(usec uint32) {
	usleep(usec)
}

//go:noescape
func writeConsole(p uintptr, n int32) int32

// write1 calls the write system call.
// It returns a non-negative number of bytes written or a negative errno value.
func write1(fd uintptr, p unsafe.Pointer, n int32) int32 {
	if fd <= 2 {
		return writeConsole(uintptr(p), n)
	}

	return -1
}

//go:noescape
func open(name *byte, mode, perm int32) int32

// return value is only set on linux to be used in osinit().
func madvise(addr unsafe.Pointer, n uintptr, flags int32) int32

// exitThread terminates the current thread, writing *wait = freeMStack when
// the stack is safe to reclaim.
//
//go:noescape
func exitThread(wait *atomic.Uint32)

//go:noescape
func setitimer(mode int32, new, old *itimerval)

//go:noescape
func sigaltstack(new, old *stackt)

func raiseproc(sig uint32)

// raise sends a signal to the calling thread.
//
// It must be nosplit because it is used by the signal handler before
// it definitely has a Go stack.
//
//go:nosplit
func raise(sig uint32) {
	kill(0, getg().m.procid, sig)
}

func kill(pid uint64, tid uint64, sig uint32)

func gettid() uint64
func getnCPU() uint32
func getPageSize() uint32

type itimerval struct {
	it_interval timeval
	it_value    timeval
}

/* m */

type mOS struct {
	preemptExtLock uint32 // preemptExtLock synchronizes preemptM with entry/exit
}

/* SIGNAL */

const (
	_NSIG = 256

	_SI_USER     = 0 /* empirically true, but not what headers say */
	_SIG_BLOCK   = 1
	_SIG_UNBLOCK = 2
	_SIG_SETMASK = 3
	_SS_DISABLE  = 4
)

// sigPerThreadSyscall is only used on linux, so we assign a bogus signal
// number.
const sigPerThreadSyscall = 1 << 31

type sigset struct{}

//go:nosplit
func runPerThreadSyscall() {
	throw("runPerThreadSyscall only valid on linux")
}

// sigFromUser reports whether the signal was sent because of a call
// to kill.
//
//go:nosplit
func (c *sigctxt) sigFromUser() bool {
	return c.sigcode() == _SI_USER
}

func nanotime1() int64

//go:noescape
func getclock_rtc(tval uintptr)

func walltime() (sec int64, nsec int32) {
	var val timeval

	getclock_rtc(uintptr(unsafe.Pointer(&val)))

	return val.tv_sec, int32(val.tv_usec)
}

const preemptMSupported = false

// suspendLock protects simultaneous SuspendThread operations from
// suspending each other.
var suspendLock mutex

func threadPreempt(tid uint64, pc uintptr, sp uintptr) int

func preemptM(mp *m) {
	if mp == getg().m {
		throw("self-preempt")
	}

	// Synchronize with external code that may try to ExitProcess.
	if !atomic.Cas(&mp.preemptExtLock, 0, 1) {
		// External code is running. Fail the preemption
		// attempt.
		mp.preemptGen.Add(1)
		return
	}

	if mp.procid == 0 {
		// The M hasn't been minit'd yet (or was just unminit'd).
		atomic.Store(&mp.preemptExtLock, 0)
		mp.preemptGen.Add(1)
		return
	}

	// Serialize thread suspension. SuspendThread is asynchronous,
	// so it's otherwise possible for two threads to suspend each
	// other and deadlock. We must hold this lock until after
	// GetThreadContext, since that blocks until the thread is
	// actually suspended.
	lock(&suspendLock)

	if threadPreempt(mp.procid, abi.FuncPCABI0(asyncPreempt), 0) != 0 {
		//thread was not in a preemptable state
		atomic.Store(&mp.preemptExtLock, 0)
		mp.preemptGen.Add(1)
		unlock(&suspendLock)
		return
	}

	unlock(&suspendLock)
	atomic.Store(&mp.preemptExtLock, 0)

	// Acknowledge the preemption.
	mp.preemptGen.Add(1)
}
