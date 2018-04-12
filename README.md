# go-strftime

[![GoDoc](https://godoc.org/github.com/imperfectgo/go-strftime?status.svg)](https://godoc.org/github.com/imperfectgo/go-strftime)
[![Build Status](https://travis-ci.org/imperfectgo/go-strftime.svg?branch=master)](https://travis-ci.org/imperfectgo/go-strftime)
[![Go Report Card](https://goreportcard.com/badge/github.com/imperfectgo/go-strftime)](https://goreportcard.com/report/github.com/imperfectgo/go-strftime)
[![Coverage](https://codecov.io/gh/imperfectgo/go-strftime/branch/master/graph/badge.svg)](https://codecov.io/gh/imperfectgo/go-strftime)

High performance C99-compatible `strftime` formatter for Go.

## Caveats

**EXPERIMENTAL** Please test before use.

## Compatibility

| Specifier |                                   Description                                    |
| :-------: | -------------------------------------------------------------------------------- |
|   `%a`    | abbreviated weekday name (Sun)                                                   |
|   `%A`    | full weekday name (Sunday)                                                       |
|   `%b`    | abbreviated month name (Sep)                                                     |
|   `%B`    | full month name (September)                                                      |
|   `%c`    | the same as time.ANSIC (%a %b %e %H:%M:%S %)                                     |
|   `%C`    | (year / 100) as number. Single digits are preceded by zero (20)                  |
|   `%d`    | day of month as number. Single digits are preceded by zero (21)                  |
|   `%D`    | equivalent to %m/%d/%y (09/21/14)                                                |
|   `%e`    | day of month as number. Single digits are preceded by a blank (21)               |
|   `%f`    | microsecond as a six digit decimal number, zero-padded on the left (001234)      |
|   `%F`    | equivalent to %Y-%m-%d (2014-09-21)                                              |
|   `%g`    | last two digits of ISO 8601 week-based year                                      |
|   `%G`    | ISO 8601 week-based year                                                         |
|   `%h`    | same as %b                                                                       |
|   `%H`    | the hour (24 hour clock) as a number. Single digits are preceded by zero (15)    |
|   `%I`    | the hour (12 hour clock) as a number. Single digits are preceded by zero (03)    |
|   `%j`    | the day of the year as a decimal number. Single digits are preced by zeros (264) |
|   `%m`    | the month as a decimal number. Single digits are preceded by a zero (09)         |
|   `%M`    | the minute as a decimal number. Single digits are preceded by a zero (32)        |
|   `%n`    | a newline (\n)                                                                   |
|   `%p`    | AM or PM as appropriate                                                          |
|   `%P`    | am or pm as appropriate                                                          |
|   `%r`    | equivalent to %I:%M:%S %p                                                        |
|   `%R`    | equivalent to %H:%M                                                              |
|   `%S`    | the second as a number. Single digits are preceded by a zero (05)                |
|   `%t`    | a tab (\t)                                                                       |
|   `%T`    | equivalent to %H:%M:%S                                                           |
|   `%u`    | weekday as a decimal number, where Monday is 1                                   |
|   `%U`    | week of the year as a decimal number (Sunday is the first day of the week)       |
|   `%V`    | ISO 8601 week of the year                                                        |
|   `%w`    | the weekday (Sunday as first day of the week) as a number. (0)                   |
|   `%W`    | week of the year as a decimal number (Monday is the first day of the week)       |
|   `%x`    | equivalent to %m/%d/%Y                                                           |
|   `%X`    | equivalent to %H:%M:%S                                                           |
|   `%y`    | year without century as a number. Single digits are preceded by zero (14)        |
|   `%Y`    | the year with century as a number (2014)                                         |
|   `%z`    | the time zone offset from UTC (-0700)                                            |
|   `%Z`    | time zone name (UTC)                                                             |

## Performance

Comparision with the standard library `time.(*Time).Format()`:

```
> go test -tags bench -bench Bench -cpu 4 -benchmem .

goos: darwin
goarch: amd64
pkg: github.com/imperfectgo/go-strftime
BenchmarkStdTimeFormat-4         5000000               356 ns/op              48 B/op          1 allocs/op
BenchmarkGoStrftime-4            5000000               347 ns/op              32 B/op          1 allocs/op
PASS
ok      github.com/imperfectgo/go-strftime      4.245s
```

Comparision with other libraries:

```
> go test -tags benchcomp -bench Bench -cpu 8 -benchmem .

goos: darwin
goarch: amd64
pkg: github.com/imperfectgo/go-strftime/benchmark
BenchmarkImperfectGo-8                   3000000               484 ns/op              64 B/op          1 allocs/op
BenchmarkTebeka-8                         300000              4161 ns/op             272 B/op         20 allocs/op
BenchmarkJehiah-8                        1000000              1719 ns/op             256 B/op         17 allocs/op
BenchmarkFastly-8                        2000000               708 ns/op              85 B/op          5 allocs/op
BenchmarkLestrrat-8                      1000000              1471 ns/op             240 B/op          3 allocs/op
BenchmarkLestrratCachedString-8          3000000               496 ns/op             128 B/op          2 allocs/op
PASS
ok      github.com/imperfectgo/go-strftime/benchmark    10.605s
```

## License

This project can be treated as a derived work of time package from golang standard library.

Licensed under the Modified (3-clause) BSD license.
