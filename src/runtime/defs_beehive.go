// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

const (
	_ENOSYS     = 0x1
	_EINTR      = 0x2
	_EACCESS    = 0x3
	_EFAULT     = 0x4
	_ENOPROC    = 0x5
	_ESIZE      = 0x6
	_EEXISTS    = 0x7
	_EEXHAUSTED = 0x8
	_EINUSE     = 0x9
	_EAGAIN     = 0xA

	_MMAP_READ  = 0x1
	_MMAP_WRITE = 0x2

	_SIGURG  = 0x1
	_SIGHUP  = 0x2
	_SIGINT  = 0x3
	_SIGPIPE = 0x4
	_SIGPROF = 0x5
	_SIGTRAP = 0x6
	_SIGUSR1 = 0x7
	_SIGQUIT = 0x8
	_SIGSEGV = 0x9
	_SIGILL  = 0xA
	_SIGFPE  = 0xB
	_SIGBUS  = 0xC
	_SIGABRT = 0xD

	_ITIMER_REAL    = 0x0
	_ITIMER_VIRTUAL = 0x1
	_ITIMER_PROF    = 0x2

	_CLOCK_REALTIME  = 0
	_CLOCK_VIRTUAL   = 1
	_CLOCK_PROF      = 2
	_CLOCK_MONOTONIC = 3
)
