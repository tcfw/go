// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build beehive

package time

var platformZoneSources []string // none: Beehive uses system calls instead

func initLocal() {
	localLoc.name = "UTC"
	return
}
