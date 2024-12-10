// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"errors"
	"internal/poll"
	"sync/atomic"
	"time"
)

type file struct {
	pfd        poll.FD
	name       string
	dirinfo    atomic.Pointer[dirInfo] // nil unless directory being read
	appendMode bool                    // whether file is opened for appending
}

func NewFile(fd uintptr, name string) *File {
	return nil
}

func (f *File) Fd() uintptr {
	return 0
}

func (f *File) Close() error {
	return errors.New("not implemented")
}

func (f *File) Stat() (FileInfo, error) {
	return nil, errors.New("not implemented")
}

func (f *File) Truncate(size int64) error {
	return errors.New("not implemented")
}

func (f *File) chmod(mode FileMode) error {
	return errors.New("not implemented")
}

func (f *File) Sync() error {
	return errors.New("not implemented")
}

func (f *File) read(b []byte) (n int, err error) {
	return 0, errors.New("not implemented")
}

func (f *File) pread(b []byte, off int64) (n int, err error) {
	return 0, errors.New("not implemented")
}

func (f *File) write(b []byte) (n int, err error) {
	return 0, errors.New("not implemented")
}

func (f *File) pwrite(b []byte, off int64) (n int, err error) {
	return 0, errors.New("not implemented")
}

func (f *File) seek(offset int64, whence int) (ret int64, err error) {
	return 0, errors.New("not implemented")
}

func (f *File) Chown(uid, gid int) error {
	return errors.New("not implemented")
}

func (f *File) Chdir() error {
	return errors.New("not implemented")
}

func (f *File) setDeadline(time.Time) error {
	return errors.New("not implemented")
}

func (f *File) setReadDeadline(time.Time) error {
	return errors.New("not implemented")
}

func (f *File) setWriteDeadline(time.Time) error {
	return errors.New("not implemented")
}

func (f *File) checkValid(op string) error {
	return errors.New("not implemented")
}

func rename(oldname, newname string) error {
	return errors.New("not implemented")
}

func readlink(name string) (string, error) {
	return "", errors.New("not implemented")
}

func tempDir() string {
	return ""
}

func epipecheck(file *File, e error) {
}

// fixLongPath is a noop on non-Windows platforms.
func fixLongPath(path string) string {
	return path
}

// See docs in file.go:Chmod.
func chmod(name string, mode FileMode) error {
	return errors.New("not implemented")
}

func ignoringEINTR(fn func() error) error {
	return fn()
}

// Remove removes the named file or directory.
// If there is an error, it will be of type *PathError.
func Remove(name string) error {
	return errors.New("not implemented")
}

func openFileNolog(name string, flag int, perm FileMode) (*File, error) {
	return nil, errors.New("not implemented")
}

func openDirNolog(name string) (*File, error) {
	return openFileNolog(name, O_RDONLY, 0)
}

// syscallMode returns the syscall-specific mode bits from Go's portable mode bits.
func syscallMode(i FileMode) (o uint32) {
	return uint32(i)
}
