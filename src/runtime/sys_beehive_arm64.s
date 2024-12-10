// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//
// System calls and other sys.stuff for arm64, Linux
//

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"
#include "cgo/abi_arm64.h"

#define SYSCALL_get_pid				1
#define SYSCALL_get_ppid			2
#define SYSCALL_get_tid				3
#define SYSCALL_pctrl				4
#define SYSCALL_set_uid				5
#define SYSCALL_set_gid				6
#define SYSCALL_get_uid				7
#define SYSCALL_get_gid				8
#define SYSCALL_get_cpu				9
#define SYSCALL_running_time		10
#define SYSCALL_proc_info			11
#define SYSCALL_get_time			15
#define SYSCALL_sched_getaffinity	20
#define SYSCALL_sched_setaffinity	21
#define SYSCALL_sched_getpriority	22
#define SYSCALL_sched_setpriority	23
#define SYSCALL_sched_yield			24
#define SYSCALL_exec				30
#define SYSCALL_clone				31
#define SYSCALL_thread_start		32
#define SYSCALL_thread_preempt		33
#define SYSCALL_kname				40
#define SYSCALL_sysinfo				41
#define SYSCALL_set_hostname		42
#define SYSCALL_get_hostname		43
#define SYSCALL_mq_open				50
#define SYSCALL_mq_close			51
#define SYSCALL_mq_ctrl				52
#define SYSCALL_mq_send				53
#define SYSCALL_mq_msend			54
#define SYSCALL_mq_recv				55
#define SYSCALL_mq_mrecv			56
#define SYSCALL_mq_notify			57
#define SYSCALL_shm_get				58
#define SYSCALL_shm_ctrl			59
#define SYSCALL_shm_attach			60
#define SYSCALL_shm_detach			61
#define SYSCALL_sem_get				62
#define SYSCALL_sem_op				63
#define SYSCALL_sem_ctl				64
#define SYSCALL_futex				65
#define SYSCALL_mem_map				80
#define SYSCALL_mem_umap			81
#define SYSCALL_mem_protect			82
#define SYSCALL_brk					83
#define SYSCALL_get_random			90
#define SYSCALL_shutdown			91
#define SYSCALL_kill				100
#define SYSCALL_exit				101
#define SYSCALL_exit_group			102
#define SYSCALL_wait				103
#define SYSCALL_sig_wait			104
#define SYSCALL_sig_action			105
#define SYSCALL_sig_pending			106
#define SYSCALL_sig_mask			107
#define SYSCALL_sig_return			108
#define SYSCALL_dev_sigmap			109
#define SYSCALL_dev_sigack			110
#define SYSCALL_sig_tstack			111
#define SYSCALL_nanosleep			120
#define SYSCALL_timer_create		121
#define SYSCALL_timer_get			122
#define SYSCALL_timer_set			123
#define SYSCALL_timer_delete		124
#define SYSCALL_timer_overrun		125
#define SYSCALL_console_read		130
#define SYSCALL_console_write		131
#define SYSCALL_trace				140
#define SYSCALL_dev_count			150
#define SYSCALL_dev_info			151
#define SYSCALL_dev_claim			152
#define SYSCALL_dev_release			153

#define FUTEX_OP_wake	1
#define FUTEX_OP_sleep	2

TEXT runtime·exit(SB),NOSPLIT|NOFRAME,$0-4
	MOVD	$SYSCALL_exit_group, R0
	MOVW	code+0(FP), R1

	SVC

	RET

// func exitThread(wait *atomic.Uint32)
TEXT runtime·exitThread(SB),NOSPLIT|NOFRAME,$0-8
	MOVD	$SYSCALL_exit, R0
	MOVW	$0, R1	// exit code
	MOVD	wait+0(FP), R2

	SVC

	JMP	0(PC)

TEXT runtime·gettid(SB),NOSPLIT,$0-8
	MOVD	$SYSCALL_get_tid, R0

	SVC

	MOVW	R0, ret+0(FP)
	RET

TEXT runtime·getpid(SB),NOSPLIT,$0-8
	MOVD	$SYSCALL_get_pid, R0

	SVC

	MOVD	R0, ret+0(FP)
	RET

TEXT runtime·sysMmap(SB),NOSPLIT,$0
	MOVD	$SYSCALL_mem_map, R0
	MOVD	$0, R1
	MOVD	addr+0(FP), R2
	MOVD	n+8(FP), R3
	MOVW	flags+16(FP), R4
	MOVD	$0, R5

	SVC

	CMN		$4095, R0
	BCC		mmapOk
	NEG		R0, R0
	MOVD	$0, p+24(FP)
	MOVD	R0, err+32(FP)
	RET
mmapOk:
	MOVD	R0, p+24(FP)
	MOVD	$0, err+32(FP)
	RET

TEXT runtime·sysMunmap(SB),NOSPLIT,$0
	MOVD	$SYSCALL_mem_umap, R0
	MOVD	addr+0(FP), R1
	MOVD	n+8(FP), R2

	SVC

	CMN		$4095, R0
	BCC		munmapOk
	MOVD	R0, 0xf0(R0)
munmapOk:
	RET

TEXT runtime·osyield(SB),NOSPLIT|NOFRAME,$0
	MOVD	$SYSCALL_sched_yield, R0
	SVC
	RET

TEXT runtime·nanotime1(SB),NOSPLIT,$0-8
	MOVD $SYSCALL_get_time, R0
	MOVD $1, R1 //monotonic

	SVC

	MOVD	R0, ret+0(FP)

	RET

TEXT runtime·getclock_rtc(SB),NOSPLIT,$0-8
	MOVD $SYSCALL_get_time, R0
	MOVD $0, R1 //RTC
	MOVD tval+0(FP), R2
	SVC
	RET

// func usleep(usec uint32)
TEXT runtime·usleep(SB),NOSPLIT,$24-4
	MOVWU	usec+0(FP), R3
	MOVD	R3, R5
	MOVW	$1000000, R4
	UDIV	R4, R3
	MOVD	R3, 8(RSP)
	MUL		R3, R4
	SUB		R4, R5
	MOVW	$1000, R4
	MUL		R4, R5
	MOVD	R5, 16(RSP)

	// nanosleep(&ts, 0)
	MOVD	$SYSCALL_nanosleep, R0
	ADD		$8, RSP, R1
	MOVD	$0, R2
	SVC
	RET

TEXT runtime·kill(SB),NOSPLIT|NOFRAME,$0-16
	MOVD	$SYSCALL_kill, R0
	MOVD	pid+0(FP), R1
	MOVD	tid+8(FP), R2
	MOVD	sig+16(FP), R3
	SVC
	RET

TEXT runtime·getnCPU(SB),NOSPLIT,$0-4
	MOVD	$SYSCALL_sysinfo, R0
	MOVD	$1, R1 //ncpus
	SVC
	MOVW	R0, ret+0(FP)
	RET

TEXT runtime·getPageSize(SB),NOSPLIT,$0-4
	MOVD	$SYSCALL_sysinfo, R0
	MOVD	$2, R1 //page_size
	SVC
	MOVW	R0, ret+0(FP)
	RET

TEXT runtime·sigaltstack(SB),NOSPLIT,$0
	MOVD	$SYSCALL_sig_tstack, R0
	MOVD	new+0(FP), R1
	MOVD	old+8(FP), R2
	SVC
	CMN	$4095, R0
	BCC	ok
	MOVD	$0, R0
	MOVD	R0, (R0)	// crash
ok:
	RET

TEXT runtime·futexwakeup(SB),NOSPLIT|NOFRAME,$0
	MOVD	$SYSCALL_futex, R0
	MOVD	addr+0(FP), R1
	MOVD	$FUTEX_OP_wake, R2
	MOVD	cnt+8(FP), R3
	MOVD	$0, R4
	MOVD	$0, R5

	SVC
	RET

TEXT runtime·futexsleep(SB),NOSPLIT|NOFRAME,$0
	//futex(*addr, FUTEX_OP_sleep, val, val2, ts_ns)
	MOVD	$SYSCALL_futex, R0
	MOVD	addr+0(FP), R1
	MOVD	$FUTEX_OP_sleep, R2
	MOVW	val+8(FP), R3
	MOVD	$0, R4
	MOVD	ns+16(FP), R5

	SVC

	RET

TEXT emptyfunc<>(SB),0,$0-0
	RET

TEXT runtime·threadInit(SB),NOSPLIT|NOFRAME,$0
	// set up g
	MOVD	m_g0(R0), g
	MOVD	R0, g_m(g)
	BL	emptyfunc<>(SB)	 // fault if stack check is wrong
	BL	runtime·mstart(SB)

	MOVD	$2, R8	// crash (not reached)
	MOVD	R8, (R8)
	RET

TEXT runtime·threadStart(SB),NOSPLIT|NOFRAME,$0-4
	MOVD	$SYSCALL_thread_start, R0
	MOVD	f+0(FP), R1
	MOVD	s+8(FP), R2
	MOVD	a+16(FP), R3

	SVC

	MOVW	R0, ret+0(FP)
	RET

TEXT runtime·writeConsole(SB),NOSPLIT|NOFRAME,$0
	MOVD	$SYSCALL_console_write, R0
	MOVD	p+0(FP), R1
	MOVD	n+8(FP), R2

	SVC

	MOVW	R0, ret+0(FP)
	RET

TEXT runtime·sigaction(SB),NOSPLIT|NOFRAME,$0
	MOVD	$SYSCALL_sig_action, R0
	MOVW	i+0(FP),R1
	MOVD	new+4(FP), R2
	MOVD	old+12(FP), R3

	SVC

	MOVW	R0, ret+0(FP)
	RET

TEXT runtime·raiseproc(SB),NOSPLIT|NOFRAME,$0
	MOVD	$SYSCALL_get_pid, R0
	SVC
	MOVD	R0, R1
	MOVD	$SYSCALL_kill, R0
	MOVW	sig+0(FP), R2
	SVC
	RET

TEXT runtime·threadPreempt(SB),NOSPLIT|NOFRAME,$0-4
	MOVD	$SYSCALL_thread_preempt, R0
	MOVD	tid+0(FP), R1
	MOVD	pc+4(FP), R2
	MOVD	sp+12(FP), R3

	SVC

	MOVW	R0, ret+0(FP)
	RET
