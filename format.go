// Copyright 2018 Timon Wong. All rights reserved.
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strftime

import (
	"time"
)

const (
	_                      = iota
	stdYield                                              // Yielded chunk
	stdLongMonth           = iota + stdNeedDate           // "January"
	stdMonth                                              // "Jan"
	stdNumMonth                                           // "1"
	stdZeroMonth                                          // "01"
	stdLongWeekDay                                        // "Monday"
	stdNumWeekDay                                         // numerical week representation (0 - Sunday ~ 6 - Saturday)
	stdWeekDay                                            // "Mon"
	stdDay                                                // "2"
	stdUnderDay                                           // "_2"
	stdZeroDay                                            // "02"
	stdHour                = iota + stdNeedClock          // "15"
	stdHour12                                             // "3"
	stdZeroHour12                                         // "03"
	stdMinute                                             // "4"
	stdZeroMinute                                         // "04"
	stdSecond                                             // "5"
	stdZeroSecond                                         // "05"
	stdLongYear            = iota + stdNeedDate           // "2006"
	stdYear                                               // "06"
	stdFirstTwoDigitYear                                  // "20"
	stdYearDay                                            // day of the year as a decimal number (range [001,366])
	stdISO8601WeekYear     = iota + stdNeedISOISO8601Week // last two digits of ISO 8601 week-based year
	stdISO8601LongWeekYear                                // ISO 8601 week-based year
	stdISO8601Week                                        // ISO 8601 week
	stdPM                  = iota + stdNeedClock          // "PM"
	stdpm                                                 // "pm"
	stdTZ                  = iota                         // "MST"
	stdNumTZ                                              // "-0700"  // always numeric
	stdFracSecond0                                        // ".0", ".00", ... , trailing zeros included
	stdFracSecond9                                        // ".9", ".99", ..., trailing zeros omitted
	stdNeedDate            = 1 << 8                       // need month, day, year
	stdNeedClock           = 2 << 8                       // need hour, minute, second
	stdNeedISOISO8601Week  = 4 << 8                       // need ISO8601 week and year
	stdArgShift            = 16                           // extra argument in high bits, above low stdArgShift
	stdMask                = 1<<stdArgShift - 1           // mask out argument
)

// Format returns a textual representation of the time value formatted
// according to C99-compatible strftime layout.
func Format(t time.Time, layout string) string {
	const bufSize = 64
	var b [bufSize]byte
	var buf = b[:0]
	buf = AppendFormat(buf, t, layout)
	return string(buf)
}

// AppendFormat is like Format but appends the textual
// representation to b and returns the extended buffer.
func AppendFormat(b []byte, t time.Time, layout string) []byte {
	var (
		name, offset, abs = locabs(&t)

		year            = -1
		month           time.Month
		day             int
		yday            int
		hour            = -1
		min             int
		sec             int
		iso8601WeekYear = -1
		iso8601Week     int
	)

	// Each iteration generates one std value.
	for layout != "" {
		prefix, std, suffix := nextStdChunk(layout)
		if prefix != "" {
			b = append(b, prefix...)
		}
		if std == 0 {
			break
		}
		layout = suffix

		// Compute year, month, day if needed.
		if year < 0 && std&stdNeedDate != 0 {
			year, month, day, yday = absDate(abs, true)
		}

		// Compute hour, minute, second if needed.
		if hour < 0 && std&stdNeedClock != 0 {
			hour, min, sec = absClock(abs)
		}

		// Compute ISO8601 week year if needed
		if iso8601WeekYear < 0 && std&stdNeedISOISO8601Week != 0 {
			iso8601WeekYear, iso8601Week = t.ISOWeek()
		}

		switch std & stdMask {
		case stdYield:
			continue
		case stdISO8601WeekYear:
			b = appendInt(b, iso8601WeekYear/100, 2)
		case stdISO8601LongWeekYear:
			b = appendInt(b, iso8601WeekYear, 4)
		case stdISO8601Week:
			b = appendInt(b, iso8601Week, 2)
		case stdYear:
			y := year
			if y < 0 {
				y = -y
			}
			b = appendInt(b, y%100, 2)
		case stdLongYear:
			b = appendInt(b, year, 4)
		case stdFirstTwoDigitYear:
			b = appendInt(b, year/100, 2)
		case stdYearDay:
			b = appendInt(b, yday, 3)
		case stdMonth:
			b = append(b, month.String()[:3]...)
		case stdLongMonth:
			m := month.String()
			b = append(b, m...)
		case stdNumMonth:
			b = appendInt(b, int(month), 0)
		case stdZeroMonth:
			b = appendInt(b, int(month), 2)
		case stdWeekDay:
			b = append(b, absWeekday(abs).String()[:3]...)
		case stdLongWeekDay:
			s := absWeekday(abs).String()
			b = append(b, s...)
		case stdNumWeekDay:
			w := int(absWeekday(abs))
			b = appendInt(b, w, 0)
		case stdDay:
			b = appendInt(b, day, 0)
		case stdUnderDay:
			if day < 10 {
				b = append(b, ' ')
			}
			b = appendInt(b, day, 0)
		case stdZeroDay:
			b = appendInt(b, day, 2)
		case stdHour:
			b = appendInt(b, hour, 2)
		case stdHour12:
			// Noon is 12PM, midnight is 12AM.
			hr := hour % 12
			if hr == 0 {
				hr = 12
			}
			b = appendInt(b, hr, 0)
		case stdZeroHour12:
			// Noon is 12PM, midnight is 12AM.
			hr := hour % 12
			if hr == 0 {
				hr = 12
			}
			b = appendInt(b, hr, 2)
		case stdMinute:
			b = appendInt(b, min, 0)
		case stdZeroMinute:
			b = appendInt(b, min, 2)
		case stdSecond:
			b = appendInt(b, sec, 0)
		case stdZeroSecond:
			b = appendInt(b, sec, 2)
		case stdPM:
			if hour >= 12 {
				b = append(b, "PM"...)
			} else {
				b = append(b, "AM"...)
			}
		case stdpm:
			if hour >= 12 {
				b = append(b, "pm"...)
			} else {
				b = append(b, "am"...)
			}
		case stdNumTZ:
			zone := offset / 60 // convert to minutes
			if zone < 0 {
				b = append(b, '-')
				zone = -zone
			} else {
				b = append(b, '+')
			}
			b = appendInt(b, zone/60, 2)
			b = appendInt(b, zone%60, 2)

		case stdTZ:
			if name != "" {
				b = append(b, name...)
				break
			}
			// No time zone known for this time, but we must print one.
			// Use the -0700 format.
			zone := offset / 60 // convert to minutes
			if zone < 0 {
				b = append(b, '-')
				zone = -zone
			} else {
				b = append(b, '+')
			}
			b = appendInt(b, zone/60, 2)
			b = appendInt(b, zone%60, 2)
		case stdFracSecond0, stdFracSecond9:
			b = formatNano(b, uint(t.Nanosecond()), std>>stdArgShift, std&stdMask == stdFracSecond9)
		}
	}
	return b
}

// nextStdChunk finds the first occurrence of a std string in
// layout and returns the text before, the std string, and the text after.
func nextStdChunk(layout string) (prefix string, std int, suffix string) {
	specPos := -1

	for i := 0; i < len(layout); i++ {
		c := int(layout[i])
		if specPos < 0 {
			if c == '%' {
				specPos = i
			}
			continue
		}

		switch c {
		case 'a': // Mon
			return layout[0:specPos], stdWeekDay, layout[i+1:]
		case 'A': // Monday
			return layout[0:specPos], stdLongWeekDay, layout[i+1:]
		case 'b', 'h': // Jan
			return layout[0:specPos], stdMonth, layout[i+1:]
		case 'B': // January
			return layout[0:specPos], stdLongMonth, layout[i+1:]
		case 'c': // "Mon Jan _2 15:04:05 2006"
			return layout[0:specPos], stdYield, "%a %b %e %H:%M:%S %Y" + layout[i+1:]
		case 'C': // 20
			return layout[0:specPos], stdFirstTwoDigitYear, layout[i+1:]
		case 'd': // 02
			return layout[0:specPos], stdZeroDay, layout[i+1:]
		case 'D': // %m/%d/%y
			return layout[0:specPos], stdYield, "%m/%d/%y" + layout[i+1:]
		case 'e': // _2
			return layout[0:specPos], stdUnderDay, layout[i+1:]
		case 'f': // fraction seconds in microseconds (Python)
			std = stdFracSecond0
			std |= 6 << stdArgShift // microseconds precision
			return layout[0:specPos], std, layout[i+1:]
		case 'F': // %Y-%m-%d
			return layout[0:specPos], stdYield, "%Y-%m-%d" + layout[i+1:]
		case 'g':
			return layout[0:specPos], stdISO8601WeekYear, layout[i+1:]
		case 'G':
			return layout[0:specPos], stdISO8601LongWeekYear, layout[i+1:]
		case 'H':
			return layout[0:specPos], stdHour, layout[i+1:]
		case 'I':
			return layout[0:specPos], stdZeroHour12, layout[i+1:]
		case 'j':
			return layout[0:specPos], stdYearDay, layout[i+1:]
		case 'm':
			return layout[0:specPos], stdZeroMonth, layout[i+1:]
		case 'M':
			return layout[0:specPos], stdZeroMinute, layout[i+1:]
		case 'n':
			return layout[0:specPos] + "\n", stdYield, layout[i+1:]
		case 'p':
			return layout[0:specPos], stdPM, layout[i+1:]
		case 'P':
			return layout[0:specPos], stdpm, layout[i+1:]
		case 'r':
		case 'R': // %H:%M"
			return layout[0:specPos], stdYield, "%H:%M" + layout[i+1:]
		case 'S':
			return layout[0:specPos], stdZeroSecond, layout[i+1:]
		case 't':
			return layout[0:specPos] + "\t", stdYield, layout[i+1:]
		case 'T': // %H:%M:%S
			return layout[0:specPos], stdYield, "%H:%M:%S" + layout[i+1:]
		//case 'u': // TODO ISO8601 weekday
		//return layout[0:specPos], stdNumWeekDay, layout[i+1:]
		case 'U':
			// TODO week of the year as a decimal number (Sunday is the first day of the week)
		case 'V':
			return layout[0:specPos], stdISO8601Week, layout[i+1:]
		case 'w':
			return layout[0:specPos], stdNumWeekDay, layout[i+1:]
		case 'W':
			// TODO: week of the year as a decimal number (Monday is the first day of the week)
		//case 'x': // locale depended, not supported
		//case 'X': // locale depended, not supported
		case 'y':
			return layout[0:specPos], stdYear, layout[i+1:]
		case 'Y':
			return layout[0:specPos], stdLongYear, layout[i+1:]
		case 'z':
			return layout[0:specPos], stdNumTZ, layout[i+1:]
		case 'Z':
			return layout[0:specPos], stdTZ, layout[i+1:]
		case '%':
		}

		specPos = -1
	}

	return layout, 0, ""
}

// appendInt appends the decimal form of x to b and returns the result.
// If the decimal form (excluding sign) is shorter than width, the result is padded with leading 0's.
// Duplicates functionality in strconv, but avoids dependency.
// Duplicated from the standard Go library.
func appendInt(b []byte, x int, width int) []byte {
	u := uint(x)
	if x < 0 {
		b = append(b, '-')
		u = uint(-x)
	}

	// Assemble decimal in reverse order.
	var buf [20]byte
	i := len(buf)
	for u >= 10 {
		i--
		q := u / 10
		buf[i] = byte('0' + u - q*10)
		u = q
	}
	i--
	buf[i] = byte('0' + u)

	// Add 0-padding.
	for w := len(buf) - i; w < width; w++ {
		b = append(b, '0')
	}

	return append(b, buf[i:]...)
}

// formatNano appends a fractional second, as nanoseconds, to b
// and returns the result.
// Duplicated from the standard Go library.
func formatNano(b []byte, nanosec uint, n int, trim bool) []byte {
	u := nanosec
	var buf [9]byte
	for start := len(buf); start > 0; {
		start--
		buf[start] = byte(u%10 + '0')
		u /= 10
	}

	if n > 9 {
		n = 9
	}
	if trim {
		for n > 0 && buf[n-1] == '0' {
			n--
		}
		if n == 0 {
			return b
		}
	}
	return append(b, buf[:n]...)
}
