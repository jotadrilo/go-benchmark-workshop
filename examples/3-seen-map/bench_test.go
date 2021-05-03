package bench

import (
	"fmt"
	"testing"
	"unsafe"
)

func BenchmarkSeenBool(b *testing.B) {
	benchmarksCases := []int{10_000, 20_000, 30_000}
	for _, bc := range benchmarksCases {
		b.Run(fmt.Sprintf("%d", bc), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				seen := make(map[string]bool)
				for n := 0; n < bc; n++ {
					k := fmt.Sprintf("%d", n)
					seen[k] = true
				}
			}
		})
	}
}

func BenchmarkSeenStruct(b *testing.B) {
	benchmarksCases := []int{10_000, 20_000, 30_000}
	for _, bc := range benchmarksCases {
		b.Run(fmt.Sprintf("%d", bc), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				seen := make(map[string]struct{})
				for n := 0; n < bc; n++ {
					k := fmt.Sprintf("%d", n)
					seen[k] = struct{}{}
				}
			}
		})
	}
}

// We will use unsafe.Sizeof to measure the hypothetical size in memory of each
// implementation:
//
// Sizeof takes an expression x of any type and returns the size in bytes of
// a hypothetical variable v as if v was declared via var v = x. The size does
// not include any memory possibly referenced by x. For instance, if x is a
// slice, Sizeof returns the size of the slice descriptor, not the size of the
// memory referenced by the slice. The return value of Sizeof is a Go constant.

func TestSeenBool(t *testing.T) {
	benchmarksCases := []int{10_000, 20_000, 30_000}
	for _, bc := range benchmarksCases {
		t.Run(fmt.Sprintf("%d", bc), func(t *testing.T) {
			seen := make(map[string]bool, bc)
			for n := 0; n < bc; n++ {
				k := fmt.Sprintf("%d", n)
				seen[k] = true
			}
			sizeMap := unsafe.Sizeof(seen)
			sizeKeys := uintptr(0)
			sizeValues := uintptr(0)
			for k, v := range seen {
				sizeKeys += unsafe.Sizeof(k)
				sizeValues += unsafe.Sizeof(v)
			}
			t.Errorf("map[string]bool    (map): %v bytes\n", sizeMap)
			t.Errorf("map[string]bool   (keys): %v bytes\n", sizeKeys)
			t.Errorf("map[string]bool (values): %v bytes\n", sizeValues)
		})
	}
}

func TestSeenStruct(t *testing.T) {
	benchmarksCases := []int{10_000, 20_000, 30_000}
	for _, bc := range benchmarksCases {
		t.Run(fmt.Sprintf("%d", bc), func(t *testing.T) {
			seen := make(map[string]struct{}, bc)
			for n := 0; n < bc; n++ {
				k := fmt.Sprintf("%d", n)
				seen[k] = struct{}{}
			}
			sizeMap := unsafe.Sizeof(seen)
			sizeKeys := uintptr(0)
			sizeValues := uintptr(0)
			for k, v := range seen {
				sizeKeys += unsafe.Sizeof(k)
				sizeValues += unsafe.Sizeof(v)
			}
			t.Errorf("map[string]struct{}    (map): %v bytes\n", sizeMap)
			t.Errorf("map[string]struct{}   (keys): %v bytes\n", sizeKeys)
			t.Errorf("map[string]struct{} (values): %v bytes\n", sizeValues)
		})
	}
}
