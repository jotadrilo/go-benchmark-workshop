package bench

import (
	"fmt"
	"testing"
	"time"
)

func MyFunction(n time.Duration) {
	time.Sleep(n * n * time.Millisecond)
}

func BenchmarkMyFunction(b *testing.B) {
	testCases := []time.Duration{1, 10, 25, 50, 100}
	for _, tc := range testCases {
		b.Run(fmt.Sprintf("%d", tc), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MyFunction(tc)
			}
		})
	}
}
