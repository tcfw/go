// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

//go:nosplit
func mmap(addr unsafe.Pointer, n uintptr, flags int32) (unsafe.Pointer, int) {
	return sysMmap(addr, n, flags)
}

//go:nosplit
//go:cgo_unsafe_args
func munmap(addr unsafe.Pointer, n uintptr) {
	sysMunmap(addr, n)
}

func nanotime1() int64
