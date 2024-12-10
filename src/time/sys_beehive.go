// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build beehive

package time

import "errors"

func open(name string) (uintptr, error) {
	return 0, errors.New("not implemented")
}

func read(fd uintptr, buf []byte) (int, error) {
	return 0, errors.New("not implemented")
}

func closefd(fd uintptr) {
}

func preadn(fd uintptr, buf []byte, off int) error {
	return errors.New("not implemented")
}
