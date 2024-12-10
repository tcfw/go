// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"unsafe"
)

func sysMmap(addr unsafe.Pointer, n uintptr, flags int32) (p uintptr, err int)

func sysMunmap(addr unsafe.Pointer, n uintptr)

// Don't split the stack as this method may be invoked without a valid G, which
// prevents us from allocating more stack.
//
//go:nosplit
func sysAllocOS(n uintptr) unsafe.Pointer {
	p, err := sysMmap(nil, n, _MMAP_READ|_MMAP_WRITE)
	if err != 0 {
		if err == _EACCESS {
			print("runtime: mmap: access denied\n")
			exit(2)
		}
		if err == _EAGAIN {
			print("runtime: mmap: too much locked memory (check 'ulimit -l').\n")
			exit(2)
		}

		return nil
	}

	return unsafe.Pointer(p)
}

func sysUnusedOS(v unsafe.Pointer, n uintptr) {
}

func sysUsedOS(v unsafe.Pointer, n uintptr) {
}

func sysHugePageOS(v unsafe.Pointer, n uintptr) {
}

func sysNoHugePageOS(v unsafe.Pointer, n uintptr) {
}

func sysHugePageCollapseOS(v unsafe.Pointer, n uintptr) {
}

// Don't split the stack as this function may be invoked without a valid G,
// which prevents us from allocating more stack.
//
//go:nosplit
func sysFreeOS(v unsafe.Pointer, n uintptr) {
	munmap(v, n)
}

func sysFaultOS(v unsafe.Pointer, n uintptr) {
	mmap(v, n, 0)
}

func sysReserveOS(v unsafe.Pointer, n uintptr) unsafe.Pointer {
	p, err := mmap(v, n, 0)
	if err != 0 {
		return nil
	}
	return p
}

func sysMapOS(v unsafe.Pointer, n uintptr) {
	_, err := mmap(v, n, _MMAP_READ|_MMAP_WRITE)
	if err == _EEXHAUSTED {
		throw("runtime: out of memory")
	}
	if err != 0 {
		print("runtime: mprotect(", v, ", ", n, ") returned ", err, "\n")
		throw("runtime: cannot map pages in arena address space")
	}
}
