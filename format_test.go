// Copied from https://github.com/awoodbeck/strftime
// Copyright (c) 2018 Adam Woodbeck
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package strftime

import (
	"fmt"
	"testing"
	"time"
)

var (
	t1 = time.Date(2018, time.July, 9, 13, 14, 15, 0, time.UTC)
	t2 = time.Date(1950, time.December, 10, 4, 45, 59, 0, time.UTC)
	t3 = time.Date(2016, time.January, 1, 13, 14, 15, 0, time.UTC)
	t4 = time.Date(2015, time.January, 1, 13, 14, 15, 0, time.UTC)

	tc = []struct {
		time     time.Time
		laylout  string
		expected string
	}{
		{time: t1, laylout: "%", expected: "%"},
		{time: t1, laylout: "%%", expected: "%"},
		{time: t1, laylout: "%Q", expected: "%Q"},
		{time: t1, laylout: "%%n", expected: "%n"},
		{time: t1, laylout: "%%t", expected: "%t"},
		{time: t1, laylout: "%n%t", expected: "\n\t"},
		{time: t1, laylout: "%a", expected: "Mon"},
		{time: t1, laylout: "%A", expected: "Monday"},
		{time: t1, laylout: "%b", expected: "Jul"},
		{time: t1, laylout: "%h", expected: "Jul"},
		{time: t1, laylout: "%B", expected: "July"},
		{time: t1, laylout: "%c", expected: "Mon Jul  9 13:14:15 2018"},
		{time: t1, laylout: "%C", expected: "20"},
		{time: t1, laylout: "%d", expected: "09"},
		{time: t1, laylout: "%D", expected: "07/09/18"},
		{time: t1, laylout: "%e", expected: " 9"},
		{time: t1, laylout: "%F", expected: "2018-07-09"},
		{time: t1, laylout: "%g", expected: "18"},
		{time: t1, laylout: "%G", expected: "2018"},
		{time: t1, laylout: "%H", expected: "13"},
		{time: t1, laylout: "%I", expected: "01"},
		{time: t1, laylout: "%j", expected: "190"},
		{time: t1, laylout: "%m", expected: "07"},
		{time: t1, laylout: "%M", expected: "14"},
		{time: t1, laylout: "%n", expected: "\n"},
		{time: t1, laylout: "%p", expected: "PM"},
		{time: t2, laylout: "%p", expected: "AM"},
		{time: t1, laylout: "%P", expected: "pm"},
		{time: t2, laylout: "%P", expected: "am"},
		{time: t1, laylout: "%r", expected: "01:14:15 PM"},
		{time: t2, laylout: "%r", expected: "04:45:59 AM"},
		{time: t1, laylout: "%R", expected: "13:14"},
		{time: t2, laylout: "%R", expected: "04:45"},
		{time: t1, laylout: "%S", expected: "15"},
		{time: t1, laylout: "%t", expected: "\t"},
		{time: t1, laylout: "%T", expected: "13:14:15"},
		{time: t2, laylout: "%T", expected: "04:45:59"},
		{time: t1, laylout: "%u", expected: "1"},
		{time: t2, laylout: "%u", expected: "7"},
		{time: t1, laylout: "%V", expected: "28"},
		{time: t3, laylout: "%V", expected: "53"}, // 3 January days in this week.
		{time: t4, laylout: "%V", expected: "01"}, // 4 January days in this week.
		{time: t1, laylout: "%w", expected: "1"},
		{time: t2, laylout: "%w", expected: "0"},
		{time: t1, laylout: "%x", expected: "07/09/2018"},
		{time: t1, laylout: "%X", expected: "13:14:15"},
		{time: t2, laylout: "%X", expected: "04:45:59"},
		{time: t1, laylout: "%y", expected: "18"},
		{time: t1, laylout: "%Y", expected: "2018"},
		{time: t1, laylout: "%z", expected: "+0000"},
		{time: t1, laylout: "%Z", expected: "UTC"},
		{time: t1, laylout: "foo", expected: "foo"},
		{time: t1, laylout: "bar%", expected: "bar%"},
		{time: t1, laylout: "%1", expected: "%1"},
		{time: t1, laylout: "%Y-%m-%dtest\n\t%Z", expected: "2018-07-09test\n\tUTC"},
	}
)

func TestFormat(t *testing.T) {

	for i := range tc {
		c := tc[i]
		t.Run(fmt.Sprintf("layout: %s", c.laylout), func(t *testing.T) {
			actual := Format(c.time, c.laylout)
			if actual != c.expected {
				t.Errorf("expected: %q; actual: %q", c.expected, actual)
			}
		})
	}
}
