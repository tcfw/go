// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

const (
	_FPE_INTDIV = 0x7
	_FPE_INTOVF = 0x8
	_FPE_FLTDIV = 0x1
	_FPE_FLTOVF = 0x2
	_FPE_FLTUND = 0x3
	_FPE_FLTRES = 0x4
	_FPE_FLTINV = 0x5
	_FPE_FLTSUB = 0x6

	_BUS_ADRALN = 0x1
	_BUS_ADRERR = 0x2
	_BUS_OBJERR = 0x3

	_SEGV_MAPERR = 0x1
	_SEGV_ACCERR = 0x2

	_SA_SIGINFO   = 0x40
	_SA_RESTART   = 0x2
	_SA_ONSTACK   = 0x1
	_SA_USERTRAMP = 0x100
	_SA_64REGSET  = 0x200
)

type siginfo struct {
	si_signo  int32
	si_errno  int32
	si_code   int32
	si_pid    int32
	si_uid    uint32
	si_status int32
	si_addr   *byte
	si_value  [8]byte
	si_band   int64
	__pad     [7]uint64
}

type regs64 struct {
	x     [29]uint64 // registers x0 to x28
	fp    uint64     // frame register, x29
	lr    uint64     // link register, x30
	sp    uint64     // stack pointer, x31
	pc    uint64     // program counter
	cpsr  uint32     // current program status register
	__pad uint32
}

type stackt struct {
	ss_sp    uintptr
	ss_flags int32
	ss_size  uintptr
}

type neonstate64 struct {
	v    [64]uint64 // actually [32]uint128
	fpsr uint32
	fpcr uint32
}

type exceptionstate64 struct {
	far uint64 // virtual fault addr
	esr uint32 // exception syndrome
	exc uint32 // number of arm exception taken
}

type mcontext64 struct {
	es exceptionstate64
	ss regs64
	ns neonstate64
}

type ucontext struct {
	uc_onstack  int32
	uc_sigmask  sigset
	uc_stack    stackt
	uc_link     *ucontext
	uc_mcsize   uint64
	uc_mcontext *mcontext64
}

type timeval struct {
	tv_sec  int64
	tv_usec int64
}

type timespec struct {
	tv_sec  int64
	tv_nsec int64
}

type sigactiont struct {
	sa_handler uintptr // a union of two pointer
	sa_mask    sigset
	sa_flags   int32
	pad_cgo_0  [4]byte
}

func (tv *timeval) set_usec(x int32) {
	tv.tv_usec = int64(x)
}
