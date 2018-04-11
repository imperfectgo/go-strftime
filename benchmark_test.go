// Copyright 2018 Timon Wong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build bench

package strftime_test

import (
	"testing"
	"time"

	"github.com/imperfectgo/go-strftime"
)

var (
	now = time.Now().UTC()
)

func BenchmarkStdTimeFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		now.Format(time.RFC3339Nano)
	}
}

func BenchmarkGoStrftime(b *testing.B) {
	const layout = "%Y-%m-%dT%H:%M:%S.%f%z"
	now := time.Now()
	for i := 0; i < b.N; i++ {
		strftime.Format(now, layout)
	}
}
