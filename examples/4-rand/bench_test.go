package bench

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	mrand "math/rand"
	"testing"
)

func BenchmarkCryptoRand(b *testing.B) {
	for _, n := range []int64{100, 1_000, 10_000, 100_000, 1_000_000} {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				v, err := crand.Int(crand.Reader, big.NewInt(n))
				if err != nil {
					b.Fatal(err)
				}
				v.Int64()
			}
		})
	}
}

func BenchmarkMathRand(b *testing.B) {
	for _, n := range []int64{100, 1_000, 10_000, 100_000, 1_000_000} {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mrand.Int63n(n)
			}
		})
	}
}
