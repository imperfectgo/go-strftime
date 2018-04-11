// +build !appengine,!js

// Copyright 2018 Timon Wong. All rights reserved.
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strftime

import (
	"time"
	_ "unsafe" // Required for go:linkname
)

//go:linkname locabs time.(*Time).locabs
//go:noescape
func locabs(t *time.Time) (name string, offset int, abs uint64)

//go:linkname absDate time.absDate
func absDate(abs uint64, full bool) (year int, month time.Month, day int, yday int)

//go:linkname absClock time.absClock
func absClock(abs uint64) (hour, min, sec int)

//go:linkname absWeekday time.absWeekday
func absWeekday(abs uint64) time.Weekday
