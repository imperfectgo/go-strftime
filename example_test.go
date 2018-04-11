// Copyright 2018 Timon Wong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strftime

import (
	"fmt"
	"time"
)

func ExampleFormat() {
	t := time.Date(2008, 9, 3, 20, 4, 26, 654321000, time.FixedZone("CST", 8*3600))

	fmt.Println(Format(t, "%c"))
	fmt.Println(Format(t, "%Y-%m-%dT%H:%M:%S.%f%z"))
	fmt.Println(Format(t, "%A is day number %w of the week."))
	fmt.Println(Format(t, "Last century was%n the %Cth century."))
	fmt.Println(Format(t, "Time zone: %Z"))

	// Output:
	// Wed Sep  3 20:04:26 2008
	// 2008-09-03T20:04:26.654321+0800
	// Wednesday is day number 3 of the week.
	// Last century was
	//  the 20th century.
	// Time zone: CST
}
