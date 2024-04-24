// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"internal/runtime/atomic"
	"unsafe"
)

func osyield()

//go:nosplit
func osyield_no_g() {
	osyield()
}

func osinit() {
	ncpu = int32(getncpu())
	if physPageSize == 0 {
		physPageSize = getPageSize()
	}
}

func setThreadCPUProfiler(hz int32) {
	setThreadCPUProfilerHz(hz)
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
	minitSignals()
}

// Called from dropm to undo the effect of an minit.
//
//go:nosplit
func unminit() {
	unminitSignals()
	getg().m.procid = 0
}

// Called from exitm, but not from drop, to undo the effect of thread-owned
// resources in minit, semacreate, or elsewhere. Do not take locks after calling this.
func mdestroy(mp *m) {}

// May run with m.p==nil, so write barriers are not allowed.
//
//go:nowritebarrierrec
func newosproc(mp *m) {

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

// write1 calls the write system call.
// It returns a non-negative number of bytes written or a negative errno value.
func write1(fd uintptr, p unsafe.Pointer, n int32) int32 {
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

func getncpu() uint32

type itimerval struct {
	it_interval timeval
	it_value    timeval
}

/* m */

type mOS struct{}

/* SIGNAL */

const (
	_NSIG        = 256
	_SI_USER     = 0 /* empirically true, but not what headers say */
	_SIG_BLOCK   = 1
	_SIG_UNBLOCK = 2
	_SIG_SETMASK = 3
	_SS_DISABLE  = 4
)

// sigPerThreadSyscall is only used on linux, so we assign a bogus signal
// number.
const sigPerThreadSyscall = 1 << 31

//go:nosplit
//go:nowritebarrierrec
func getsig(i uint32) uintptr

//go:nosplit
//go:nowritebarrierrec
func setsigstack(i uint32)

// setSignalstackSP sets the ss_sp field of a stackt.
//
//go:nosplit
func setSignalstackSP(s *stackt, sp uintptr) {
	*(*uintptr)(unsafe.Pointer(&s.ss_sp)) = sp
}

//go:nosplit
//go:nowritebarrierrec
func setsig(i uint32, fn uintptr) {}

//go:nosplit
//go:nowritebarrierrec
func sigaddset(mask *sigset, i int) {}

func sigdelset(mask *sigset, i int) {}

//go:nosplit
//go:nowritebarrierrec
func sigprocmask(how int32, new, old *sigset) {}

//go:nosplit
func (c *sigctxt) fixsigcode(sig uint32) {}

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

func getPageSize() uintptr {
	return 0
}

func walltime() (sec int64, nsec int32) {
	return 0, 0
}
