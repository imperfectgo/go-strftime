// Copyright 2018 Timon Wong. All rights reserved.
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strftime

import (
	"testing"
	"time"
)

func BenchmarkStdTimeFormat(b *testing.B) {
	now := time.Now()
	for i := 0; i < b.N; i++ {
		now.Format(time.RFC3339Nano)
	}
}

func BenchmarkGoStrftime(b *testing.B) {
	const layout = "%Y-%m-%dT%H:%M:%S.%f%z"
	now := time.Now()
	for i := 0; i < b.N; i++ {
		Format(now, layout)
	}
}
