package bench

import (
	"fmt"
	"testing"
	"time"
)

func MyFunctionV1(n time.Duration) {
	time.Sleep(100 * time.Millisecond)
}

func BenchmarkMyFunctionV1(b *testing.B) {
	testCases := []time.Duration{1, 10, 25, 50, 100}
	for _, tc := range testCases {
		b.Run(fmt.Sprintf("%d", tc), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MyFunctionV1(tc)
			}
		})
	}
}

func MyFunctionV2(n time.Duration) {
	time.Sleep(n * n * time.Millisecond)
}

func BenchmarkMyFunctionV2(b *testing.B) {
	testCases := []time.Duration{1, 10, 25, 50, 100}
	for _, tc := range testCases {
		b.Run(fmt.Sprintf("%d", tc), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MyFunctionV2(tc)
			}
		})
	}
}
