// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT _rt0_arm64_beehive(SB),NOSPLIT|NOFRAME,$0
	MOVD	$runtime·rt0_go(SB), R2
	BL	(R2)
exit:
	MOVD	$101, R0 // sys_exit
	MOVD	$0, R1
	SVC
	B	exit

// external linking entry point.
TEXT main(SB),NOSPLIT|NOFRAME,$0
	JMP	_rt0_arm64_beehive(SB)
