// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

//
// System call support for ARM64, FreeBSD
//

#define SYS_syscall	0

// func Syscall(trap uintptr, a1, a2, a3 uintptr) (r1, r2, err uintptr)
TEXT ·Syscall(SB),NOSPLIT,$0-56
	BL	runtime·entersyscall<ABIInternal>(SB)
	MOVD	trap+0(FP), R0	// syscall entry
	MOVD	a1+8(FP), R1
	MOVD	a2+16(FP), R2
	MOVD	a3+24(FP), R3
	SVC	$SYS_syscall
	BCC	ok
	MOVD	$-1, R1
	MOVD	R1, r1+32(FP)
	MOVD	ZR, r2+40(FP)
	MOVD	R0, err+48(FP)
	BL	runtime·exitsyscall<ABIInternal>(SB)
	RET
ok:
	MOVD	R0, r1+32(FP)
	MOVD	R1, r2+40(FP)
	MOVD	ZR, err+48(FP)
	BL	runtime·exitsyscall<ABIInternal>(SB)
	RET

// func RawSyscall(trap uintptr, a1, a2, a3 uintptr) (r1, r2, err uintptr)
TEXT ·RawSyscall(SB),NOSPLIT,$0-56
	MOVD	trap+0(FP), R0	// syscall entry
	MOVD	a1+8(FP), R1
	MOVD	a2+16(FP), R2
	MOVD	a3+24(FP), R3
	SVC	$SYS_syscall
	BCC	ok
	MOVD	$-1, R1
	MOVD	R1, r1+32(FP)
	MOVD	ZR, r2+40(FP)
	MOVD	R0, err+48(FP)
	RET
ok:
	MOVD	R0, r1+32(FP)
	MOVD	R1, r2+40(FP)
	MOVD	ZR, err+48(FP)
	RET

// func Syscall6(trap uintptr, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr)
TEXT ·Syscall6(SB),NOSPLIT,$0-80
	BL	runtime·entersyscall<ABIInternal>(SB)
	MOVD	trap+0(FP), R0	// syscall entry
	MOVD	a1+8(FP), R1
	MOVD	a2+16(FP), R2
	MOVD	a3+24(FP), R3
	MOVD	a4+32(FP), R4
	MOVD	a5+40(FP), R5
	MOVD	a6+48(FP), R6
	SVC	$SYS_syscall
	BCC	ok
	MOVD	$-1, R1
	MOVD	R1, r1+56(FP)
	MOVD	ZR, r2+64(FP)
	MOVD	R0, err+72(FP)
	BL	runtime·exitsyscall<ABIInternal>(SB)
	RET
ok:
	MOVD	R0, r1+56(FP)
	MOVD	R1, r2+64(FP)
	MOVD	ZR, err+72(FP)
	BL	runtime·exitsyscall<ABIInternal>(SB)
	RET

// func RawSyscall6(trap uintptr, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr)
TEXT ·RawSyscall6(SB),NOSPLIT,$0-80
	MOVD	trap+0(FP), R0	// syscall entry
	MOVD	a1+8(FP), R1
	MOVD	a2+16(FP), R2
	MOVD	a3+24(FP), R3
	MOVD	a4+32(FP), R4
	MOVD	a5+40(FP), R5
	MOVD	a6+48(FP), R6
	SVC	$SYS_syscall
	BCC	ok
	MOVD	$-1, R1
	MOVD	R1, r1+56(FP)
	MOVD	ZR, r2+64(FP)
	MOVD	R0, err+72(FP)
	RET
ok:
	MOVD	R0, r1+56(FP)
	MOVD	R1, r2+64(FP)
	MOVD	ZR, err+72(FP)
	RET
