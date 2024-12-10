// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Beehive system calls.
// This file is compiled as ordinary Go code,
// but it is also input to mksyscall,
// which parses the //sys lines and generates system call stubs.
// Note that sometimes we use a lowercase //sys name and
// wrap it in our own nicer implementation.

package syscall

import (
	errorspkg "errors"
	"internal/itoa"
	"internal/oserror"
)

const (
	Stdin  = 0
	Stdout = 1
	Stderr = 2

	O_RDONLY = 1 << 0
	O_WRONLY = 1 << 1
	O_RDWR   = 0
	O_APPEND = 1 << 2
	O_CREAT  = 1 << 3
	O_EXCL   = 1 << 4
	O_SYNC   = 0
	O_TRUNC  = 1 << 5
)

// A Signal is a number describing a process signal.
// It implements the [os.Signal] interface.
type Signal uint8

const (
	SIGINT Signal = iota + 1
	SIGKILL
)

func (s Signal) Signal() {}

func (s Signal) String() string {
	return "todo"
}

func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)

type Errno uintptr

func (e Errno) Error() string {
	if 0 <= int(e) && int(e) < len(errors) {
		s := errors[e]
		if s != "" {
			return s
		}
	}
	return "errno " + itoa.Itoa(int(e))
}

func (e Errno) Is(target error) bool {
	switch target {
	case oserror.ErrPermission:
		return e == EACCESS
	case oserror.ErrExist:
		return e == EEXISTS
	case oserror.ErrNotExist:
		return e == ENOENT
	case errorspkg.ErrUnsupported:
		return e == ENOSYS
	}
	return false
}

func (e Errno) Temporary() bool {
	return e == EINTR || e.Timeout()
}

func (e Errno) Timeout() bool {
	return e == EAGAIN
}

type Stat_t struct {
	Dev        uint32
	Pad0       [3]int32
	Ino        uint64
	Mode       uint32
	Nlink      uint32
	Uid        uint32
	Gid        uint32
	Rdev       uint32
	Pad1       [3]uint32
	Size       int64
	Atime      uint32
	Atime_nsec uint32
	Mtime      uint32
	Mtime_nsec uint32
	Ctime      uint32
	Ctime_nsec uint32
	Blksize    uint32
	Pad2       uint32
	Blocks     int64
}

func Mkdir(path string, mode uint32) error {
	return errorspkg.New("not implemented")
}

const ImplementsGetwd = false

func Getwd() (string, error) {
	return "", errorspkg.New("not implemented")
}

func Chdir(path string) error {
	return errorspkg.New("not implemented")
}

func Geteuid() int {
	return Getuid()
}

func Getegid() int {
	return Getgid()
}

func Getgroups() (gids []int, err error) {
	return nil, errorspkg.New("not implemented")
}

const (
	EISDIR Errno = 0
	ENAMETOOLONG
	ENOTDIR
)

//sysnb	Getpid() (pid int)
//sysnb	Getppid() (ppid int)
//sysnb	Gettid() (tid int)
//sysnb	Getuid() (uid int)
//sysnb	Getgid() (gid int)
