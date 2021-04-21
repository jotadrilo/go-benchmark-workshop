package bench

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func gen(n int) []int {
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, rand.Intn(1e9))
	}
	return s
}

// Sort slice using Golang sort implementation
func SortInts(s []int) {
	sort.Ints(s)
}

func BenchmarkSortInts(b *testing.B) {
	testCases := []int{10, 100, 1000, 10000, 100000}
	for _, tc := range testCases {
		b.Run(fmt.Sprintf("%d", tc), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Stop benchmark timer while preparing the input
				b.StopTimer()
				s := gen(tc)
				b.StartTimer()

				SortInts(s)
			}
		})
	}
}
