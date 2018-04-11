# go-strftime

[![GoDoc](https://godoc.org/github.com/imperfectgo/go-strftime?status.svg)](https://godoc.org/github.com/imperfectgo/go-strftime) 
[![Build Status](https://travis-ci.org/imperfectgo/go-strftime.svg?branch=master)](https://travis-ci.org/imperfectgo/go-strftime)
[![Go Report Card](https://goreportcard.com/badge/github.com/imperfectgo/go-strftime)](https://goreportcard.com/report/github.com/imperfectgo/go-strftime)
[![Coverage](https://codecov.io/gh/imperfectgo/go-strftime/branch/master/graph/badge.svg)](https://codecov.io/gh/imperfectgo/go-strftime)

High performance C99-compatible `strftime` formatter for Go.

## Caveats

**EXPERIMENTAL** Please test before use.

## Performance

Comparision with the standard library `time.(*Time).Format()`:

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

Comparision with other libraries:

```
> go test -bench Bench -cpu 8 -benchmem ./benchmark

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
