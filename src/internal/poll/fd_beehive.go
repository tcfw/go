// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package poll

import "errors"

// FD is a file descriptor. The os packages use this type as a
// field of a larger type representing an OS file.
type FD struct {
	// Lock sysfd and serialize access to Read and Write methods.
	fdmu fdMutex

	// System file descriptor. Immutable until Close.
	Sysfd int

	// Whether this is a file
	isFile bool
}

// Destroy closes the file descriptor. This is called when there are
// no remaining references.
func (fd *FD) destroy() error {
	return errors.New("not implemented")
}

// RawControl invokes the user-defined function f for a non-IO
// operation.
func (fd *FD) RawControl(f func(uintptr)) error {
	if err := fd.incref(); err != nil {
		return err
	}
	defer fd.decref()
	f(uintptr(fd.Sysfd))
	return nil
}

// RawRead invokes the user-defined function f for a read operation.
func (fd *FD) RawRead(f func(uintptr) bool) error {
	return errors.New("not implemented")
}

// RawWrite invokes the user-defined function f for a write operation.
func (fd *FD) RawWrite(f func(uintptr) bool) error {
	return errors.New("not implemented")
}
