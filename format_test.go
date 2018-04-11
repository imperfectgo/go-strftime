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
// FITNESS FOR A PARTICULAR PURPOSE AND NON INFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package strftime

import (
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
		layout   string
		expected string
	}{
		{time: t1, layout: "%", expected: "%"},
		{time: t1, layout: "%%", expected: "%"},
		{time: t1, layout: "%Q", expected: "%Q"},
		{time: t1, layout: "%%n", expected: "%n"},
		{time: t1, layout: "%%t", expected: "%t"},
		{time: t1, layout: "%n%t", expected: "\n\t"},
		{time: t1, layout: "%a", expected: "Mon"},
		{time: t1, layout: "%A", expected: "Monday"},
		{time: t1, layout: "%b", expected: "Jul"},
		{time: t1, layout: "%h", expected: "Jul"},
		{time: t1, layout: "%B", expected: "July"},
		{time: t1, layout: "%c", expected: "Mon Jul  9 13:14:15 2018"},
		{time: t1, layout: "%C", expected: "20"},
		{time: t1, layout: "%d", expected: "09"},
		{time: t1, layout: "%D", expected: "07/09/18"},
		{time: t1, layout: "%e", expected: " 9"},
		{time: t1, layout: "%F", expected: "2018-07-09"},
		{time: t1, layout: "%g", expected: "18"},
		{time: t1, layout: "%G", expected: "2018"},
		{time: t1, layout: "%H", expected: "13"},
		{time: t1, layout: "%I", expected: "01"},
		{time: t1, layout: "%j", expected: "190"},
		{time: t1, layout: "%m", expected: "07"},
		{time: t1, layout: "%M", expected: "14"},
		{time: t1, layout: "%n", expected: "\n"},
		{time: t1, layout: "%p", expected: "PM"},
		{time: t2, layout: "%p", expected: "AM"},
		{time: t1, layout: "%P", expected: "pm"},
		{time: t2, layout: "%P", expected: "am"},
		{time: t1, layout: "%r", expected: "01:14:15 PM"},
		{time: t2, layout: "%r", expected: "04:45:59 AM"},
		{time: t1, layout: "%R", expected: "13:14"},
		{time: t2, layout: "%R", expected: "04:45"},
		{time: t1, layout: "%S", expected: "15"},
		{time: t1, layout: "%t", expected: "\t"},
		{time: t1, layout: "%T", expected: "13:14:15"},
		{time: t2, layout: "%T", expected: "04:45:59"},
		{time: t1, layout: "%u", expected: "1"},
		{time: t2, layout: "%u", expected: "7"},
		{time: t1, layout: "%V", expected: "28"},
		{time: t3, layout: "%V", expected: "53"}, // 3 January days in this week.
		{time: t4, layout: "%V", expected: "01"}, // 4 January days in this week.
		{time: t1, layout: "%w", expected: "1"},
		{time: t2, layout: "%w", expected: "0"},
		{time: t1, layout: "%x", expected: "07/09/2018"},
		{time: t1, layout: "%X", expected: "13:14:15"},
		{time: t2, layout: "%X", expected: "04:45:59"},
		{time: t1, layout: "%y", expected: "18"},
		{time: t1, layout: "%Y", expected: "2018"},
		{time: t1, layout: "%z", expected: "+0000"},
		{time: t1, layout: "%Z", expected: "UTC"},
		{time: t1, layout: "foo", expected: "foo"},
		{time: t1, layout: "bar%", expected: "bar%"},
		{time: t1, layout: "%1", expected: "%1"},
		{time: t1, layout: "%U %W", expected: "27 28"},
		{time: t1, layout: "%Y-%m-%dtest\n\t%Z", expected: "2018-07-09test\n\tUTC"},
	}
)

func TestFormat(t *testing.T) {

	for i := range tc {
		c := tc[i]
		actual := Format(c.time, c.layout)
		if actual != c.expected {
			t.Errorf("Test layout `%s`: expected: %q; actual: %q", c.layout, c.expected, actual)
		}
	}
}
