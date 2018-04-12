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
	stdNop                                                // Yielded chunk
	stdLongMonth           = iota + stdNeedDate           // "January"
	stdMonth                                              // "Jan"
	stdNumMonth                                           // "1"
	stdZeroMonth                                          // "01"
	stdLongWeekDay                                        // "Monday"
	stdZeroBasedNumWeekDay                                // numerical week representation (0 - Sunday ~ 6 - Saturday)
	stdNumWeekDay                                         // numerical week representation (1 - Monday ~ 7- Sunday)
	stdWeekDay                                            // "Mon"
	stdWeekOfYear                                         // week of the year (Sunday first)
	stdMonFirstWeekOfYear                                 // week of the year (Monday first)
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
	stdYearDay                                            // day of the year (range [001,366])
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
//
// List of accepted specifiers:
//  %a  abbreviated weekday name (Sun)
//  %A  full weekday name (Sunday)
//  %b  abbreviated month name (Sep)
//  %B  full month name (September)
//  %c  the same as time.ANSIC (%a %b %e %H:%M:%S %)
//  %C  (year / 100) as number. Single digits are preceded by zero (20)
//  %d  day of month as number. Single digits are preceded by zero (21)
//  %D  equivalent to %m/%d/%y (09/21/14)
//  %e  day of month as number. Single digits are preceded by a blank (21)
//  %f  microsecond as a six digit decimal number, zero-padded on the left (001234)
//  %F  equivalent to %Y-%m-%d (2014-09-21)
//  %g  last two digits of ISO 8601 week-based year
//  %G  ISO 8601 week-based year
//  %h  same as %b
//  %H  the hour (24 hour clock) as a number. Single digits are preceded by zero (15)
//  %I  the hour (12 hour clock) as a number. Single digits are preceded by zero (03)
//  %j  the day of the year as a decimal number. Single digits are preced by zeros (264)
//  %m  the month as a decimal number. Single digits are preceded by a zero (09)
//  %M  the minute as a decimal number. Single digits are preceded by a zero (32)
//  %n  a newline (\n)
//  %p  AM or PM as appropriate
//  %P  am or pm as appropriate
//  %r  equivalent to %I:%M:%S %p
//  %R  equivalent to %H:%M
//  %S  the second as a number. Single digits are preceded by a zero (05)
//  %t  a tab (\t)
//  %T  equivalent to %H:%M:%S
//  %u  weekday as a decimal number, where Monday is 1
//  %U  week of the year as a decimal number (Sunday is the first day of the week)
//  %V  ISO 8601 week of the year
//  %w  the weekday (Sunday as first day of the week) as a number. (0)
//  %W  week of the year as a decimal number (Monday is the first day of the week)
//  %x  equivalent to %m/%d/%Y
//  %X  equivalent to %H:%M:%S
//  %y  year without century as a number. Single digits are preceded by zero (14)
//  %Y  the year with century as a number (2014)
//  %z  the time zone offset from UTC (-0700)
//  %Z  time zone name (UTC)
func Format(t time.Time, layout string) string {
	const bufSize = 64
	var b [bufSize]byte
	buf := AppendFormat(b[:0], t, layout)
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
		case stdNop:
			continue
		case stdISO8601WeekYear:
			b = appendInt(b, iso8601WeekYear%100, 2)
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
			b = appendInt(b, yday+1, 3)
		case stdMonth:
			b = append(b, month.String()[:3]...)
		case stdLongMonth:
			m := month.String()
			b = append(b, m...)
		//case stdNumMonth:
		//	b = appendInt(b, int(month), 0)
		case stdZeroMonth:
			b = appendInt(b, int(month), 2)
		case stdWeekDay:
			b = append(b, absWeekday(abs).String()[:3]...)
		case stdLongWeekDay:
			s := absWeekday(abs).String()
			b = append(b, s...)
		case stdZeroBasedNumWeekDay:
			w := int(absWeekday(abs))
			b = appendInt(b, w, 0)
		case stdNumWeekDay:
			w := int(absWeekday(abs))
			if w == 0 {
				w = 7
			}
			b = appendInt(b, w, 0)
		case stdWeekOfYear, stdMonFirstWeekOfYear:
			w := int(absWeekday(abs))
			n := w - (std - stdWeekOfYear)
			if n < 0 {
				n = 7
			}
			n = ((yday - n) / 7) + 1
			b = appendInt(b, n, 2)
		//case stdDay:
		//	b = appendInt(b, day, 0)
		case stdUnderDay:
			if day < 10 {
				b = append(b, ' ')
			}
			b = appendInt(b, day, 0)
		case stdZeroDay:
			b = appendInt(b, day, 2)
		case stdHour:
			b = appendInt(b, hour, 2)
		//case stdHour12:
		//	// Noon is 12PM, midnight is 12AM.
		//	hr := hour % 12
		//	if hr == 0 {
		//		hr = 12
		//	}
		//	b = appendInt(b, hr, 0)
		case stdZeroHour12:
			// Noon is 12PM, midnight is 12AM.
			hr := hour % 12
			if hr == 0 {
				hr = 12
			}
			b = appendInt(b, hr, 2)
		//case stdMinute:
		//	b = appendInt(b, min, 0)
		case stdZeroMinute:
			b = appendInt(b, min, 2)
		//case stdSecond:
		//	b = appendInt(b, sec, 0)
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
		case stdFracSecond0, stdFracSecond9:
			b = formatNano(b, uint(t.Nanosecond()), std>>stdArgShift, std&stdMask == stdFracSecond9)
		}
	}
	return b
}

// nextStdChunk finds the first occurrence of a std string in
// layout and returns the text before, the std string, and the text after.
func nextStdChunk(layout string) (prefix string, std int, suffix string) {
	for i := 0; i < len(layout); i++ {
		j := i + 1
		if int(layout[i]) == '%' && len(layout) > j {
			spec := int(layout[j])
			switch spec {
			case 'a': // Mon
				return layout[0:i], stdWeekDay, layout[j+1:]
			case 'A': // Monday
				return layout[0:i], stdLongWeekDay, layout[j+1:]
			case 'b', 'h': // Jan
				return layout[0:i], stdMonth, layout[j+1:]
			case 'B': // January
				return layout[0:i], stdLongMonth, layout[j+1:]
			case 'c': // "Mon Jan _2 15:04:05 2006" (assumes "C" locale)
				return layout[0:i], stdNop, "%a %b %e %H:%M:%S %Y" + layout[j+1:]
			case 'C': // 20
				return layout[0:i], stdFirstTwoDigitYear, layout[j+1:]
			case 'd': // 02
				return layout[0:i], stdZeroDay, layout[j+1:]
			case 'D': // %m/%d/%y
				return layout[0:i], stdNop, "%m/%d/%y" + layout[j+1:]
			case 'e': // _2
				return layout[0:i], stdUnderDay, layout[j+1:]
			case 'f': // fraction seconds in microseconds (Python)
				std = stdFracSecond0
				std |= 6 << stdArgShift // microseconds precision
				return layout[0:i], std, layout[j+1:]
			case 'F': // %Y-%m-%d
				return layout[0:i], stdNop, "%Y-%m-%d" + layout[j+1:]
			case 'g':
				return layout[0:i], stdISO8601WeekYear, layout[j+1:]
			case 'G':
				return layout[0:i], stdISO8601LongWeekYear, layout[j+1:]
			case 'H':
				return layout[0:i], stdHour, layout[j+1:]
			case 'I':
				return layout[0:i], stdZeroHour12, layout[j+1:]
			case 'j':
				return layout[0:i], stdYearDay, layout[j+1:]
			case 'm':
				return layout[0:i], stdZeroMonth, layout[j+1:]
			case 'M':
				return layout[0:i], stdZeroMinute, layout[j+1:]
			case 'n':
				return layout[0:i] + "\n", stdNop, layout[j+1:]
			case 'p':
				return layout[0:i], stdPM, layout[j+1:]
			case 'P':
				return layout[0:i], stdpm, layout[j+1:]
			case 'r':
				return layout[0:i], stdNop, "%I:%M:%S %p" + layout[j+1:]
			case 'R': // %H:%M"
				return layout[0:i], stdNop, "%H:%M" + layout[j+1:]
			case 'S':
				return layout[0:i], stdZeroSecond, layout[j+1:]
			case 't':
				return layout[0:i] + "\t", stdNop, layout[j+1:]
			case 'T': // %H:%M:%S
				return layout[0:i], stdNop, "%H:%M:%S" + layout[j+1:]
			case 'u': // weekday as a decimal number, where Monday is 1
				return layout[0:i], stdNumWeekDay, layout[j+1:]
			case 'U': // week of the year as a decimal number (Sunday is the first day of the week)
				return layout[0:i], stdWeekOfYear, layout[j+1:]
			case 'V':
				return layout[0:i], stdISO8601Week, layout[j+1:]
			case 'w':
				return layout[0:i], stdZeroBasedNumWeekDay, layout[j+1:]
			case 'W': // week of the year as a decimal number (Monday is the first day of the week)
				return layout[0:i], stdMonFirstWeekOfYear, layout[j+1:]
			case 'x': // locale depended date representation (assumes "C" locale)
				return layout[0:i], stdNop, "%m/%d/%Y" + layout[j+1:]
			case 'X': // locale depended time representation (assumes "C" locale)
				return layout[0:i], stdNop, "%H:%M:%S" + layout[j+1:]
			case 'y':
				return layout[0:i], stdYear, layout[j+1:]
			case 'Y':
				return layout[0:i], stdLongYear, layout[j+1:]
			case 'z':
				return layout[0:i], stdNumTZ, layout[j+1:]
			case 'Z':
				return layout[0:i], stdTZ, layout[j+1:]
			case '%':
				return layout[0:i] + "%", stdNop, layout[j+1:]
			}
		}
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
