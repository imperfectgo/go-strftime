# go-strftime

High performance C99-compatible `strftime` formatter for Go.

**EXPERIMENTAL** Please DO NOT USE for now.

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

See [License](./LICENSE) file.
