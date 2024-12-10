// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Following are not implemented.

func initsig(preinit bool) {
}

func sigenable(sig uint32) {
}

func sigdisable(sig uint32) {
}

func sigignore(sig uint32) {
}

func signame(sig uint32) string {
	return ""
}

// gsignalStack is unused on beehive
type gsignalStack struct{}

//go:nosplit
func crash() {
	throw("crash")
}

//go:nosplit
func msigrestore(sigmask sigset) {
}

//go:nosplit
func sigsave(p *sigset) {
}

//go:nosplit
func sigblock(exiting bool) {
}

//go:nosplit
//go:nowritebarrierrec
func clearSignalHandlers() {
}

func sigpanic() {
	throw("fault")
}
