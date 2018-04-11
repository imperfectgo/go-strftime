// Copyright 2018 Timon Wong. All rights reserved.
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strftime

import (
	"testing"
	"time"
)

const layout = "%Y-%m-%dT%H:%M:%S.%f%z"

func TestFormat(t *testing.T) {
	now := time.Now()
	t.Log(now.Format(time.RFC3339Nano))
	t.Logf("Format %s -> %s", layout, Format(now, layout))
	t.Logf("Format %s -> %s", "%f", Format(now, "%f"))
	t.Logf("Format %s -> %s", "%c", Format(now, "%c"))
	t.Logf("Format %s -> %s", "%R", Format(now, "%R"))
	t.Logf("Format %s -> %s", "%T", Format(now, "%T"))
	t.Logf("Format %s -> %s", "%j", Format(now, "%j"))
}
