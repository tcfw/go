// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

var sigtable = [...]sigTabT{}

var sigset_all = sigset{__bits: [4]uint32{^uint32(0), ^uint32(0), ^uint32(0), ^uint32(0)}}
