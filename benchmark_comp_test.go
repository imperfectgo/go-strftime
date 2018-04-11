// Copyright 2018 Timon Wong. All rights reserved.

// +build benchcomp

package strftime_test

import (
	"testing"
	"time"

	fastly "github.com/fastly/go-utils/strftime"
	imperfectgo "github.com/imperfectgo/go-strftime"
	jehiah "github.com/jehiah/go-strftime"
	lestrrat "github.com/lestrrat-go/strftime"
	tebeka "github.com/tebeka/strftime"
)

const (
	benchfmt = `%A %a %B %b %d %H %I %M %m %p %S %Y %y %Z`
)

var (
	now = time.Now().UTC()
)

func BenchmarkImperfectGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		imperfectgo.Format(now, benchfmt)
	}
}

func BenchmarkTebeka(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tebeka.Format(benchfmt, now)
	}
}

func BenchmarkJehiah(b *testing.B) {
	// Grr, uses byte slices, and does it faster, but with more allocs
	for i := 0; i < b.N; i++ {
		jehiah.Format(benchfmt, now)
	}
}

func BenchmarkFastly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fastly.Strftime(benchfmt, now)
	}
}

func BenchmarkLestrrat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lestrrat.Format(benchfmt, now)
	}
}

func BenchmarkLestrratCachedString(b *testing.B) {
	f, _ := lestrrat.New(benchfmt)
	// This benchmark does not take into effect the compilation time
	for i := 0; i < b.N; i++ {
		f.FormatString(now)
	}
}
