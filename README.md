# go-strftime

[![GoDoc](https://godoc.org/github.com/imperfectgo/go-strftime?status.svg)](https://godoc.org/github.com/imperfectgo/go-strftime) 
[![Build Status](https://travis-ci.org/imperfectgo/go-strftime.svg?branch=master)](https://travis-ci.org/imperfectgo/go-strftime)
[![Go Report Card](https://goreportcard.com/badge/github.com/imperfectgo/go-strftime)](https://goreportcard.com/report/github.com/imperfectgo/go-strftime)
[![Coverage](https://codecov.io/gh/imperfectgo/go-strftime/branch/master/graph/badge.svg)](https://codecov.io/gh/imperfectgo/go-strftime)

High performance C99-compatible `strftime` formatter for Go.

## Caveats

**EXPERIMENTAL** Please test before use.

## Performance

Just for reference :P

```
> go test -bench Bench -cpu 4 -benchmem .

goos: darwin
goarch: amd64
pkg: github.com/imperfectgo/go-strftime
BenchmarkStdTimeFormat-4         5000000               356 ns/op              48 B/op          1 allocs/op
BenchmarkGoStrftime-4            5000000               347 ns/op              32 B/op          1 allocs/op
PASS
ok      github.com/imperfectgo/go-strftime      4.245s
```

## License

This project can be treated as a derived work of time package from golang standard library.
Licensed under the Modified (3-clause) BSD license.
