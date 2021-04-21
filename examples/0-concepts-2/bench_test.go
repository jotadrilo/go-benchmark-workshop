package bench

import (
	"testing"
	"time"
)

func MyFunction(n time.Duration) {
	time.Sleep(n * n * time.Millisecond)
}

func BenchmarkMyFunction_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyFunction(1)
	}
}
