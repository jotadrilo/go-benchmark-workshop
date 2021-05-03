package bench

import (
	"testing"
	"time"
)

func MyFunctionV1(n time.Duration) {
	time.Sleep(100 * time.Millisecond)
}

func BenchmarkMyFunctionV1_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyFunctionV1(1)
	}
}

func MyFunctionV2(n time.Duration) {
	time.Sleep(n * n * time.Millisecond)
}

func BenchmarkMyFunctionV2_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyFunctionV2(1)
	}
}
