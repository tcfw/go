// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"errors"
	"syscall"
	"time"
)

// The only signal values guaranteed to be present in the os package on all
// systems are os.Interrupt (send the process an interrupt) and os.Kill (force
// the process to exit). On Windows, sending os.Interrupt to a process with
// os.Process.Signal is not implemented; it will return an error instead of
// sending a signal.
var (
	Interrupt Signal = syscall.SIGINT
	Kill      Signal = syscall.SIGKILL
)

// ProcessState stores information about a process, as reported by Wait.
type ProcessState struct {
	pid int // The process's id.
}

func (p *ProcessState) userTime() time.Duration {
	return 0
}

func (p *ProcessState) systemTime() time.Duration {
	return 0
}

func (p *ProcessState) Pid() int {
	return 0
}

func (p *ProcessState) exited() bool {
	return false
}

func (p *ProcessState) success() bool {
	return false
}

func (p *ProcessState) sys() any {
	return nil
}

func (p *ProcessState) sysUsage() any {
	return nil
}

func startProcess(name string, argv []string, attr *ProcAttr) (p *Process, err error) {
	return nil, errors.New("not implemented")
}

func findProcess(pid int) (p *Process, err error) {
	// NOOP for unix.
	return newProcess(pid, 0), nil
}

func (p *Process) kill() error {
	return p.Signal(Kill)
}

func (p *Process) wait() (ps *ProcessState, err error) {
	return nil, errors.New("not implemented")
}

func (p *Process) signal(sig Signal) error {
	return errors.New("not implemented")
}

func (p *Process) release() error {
	return errors.New("not implemented")
}
