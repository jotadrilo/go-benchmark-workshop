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

// Sort slice using quicksort implementation
func SortIntsV1(s []int) {
	quicksort(s, 0, len(s)-1)
}

func partition(s []int, begin int, end int) int {
	pivot := end
	counter := begin

	for i := begin; i < end; i++ {
		if s[i] < s[pivot] {
			temp := s[counter]
			s[counter] = s[i]
			s[i] = temp
			counter++
		}
	}

	temp := s[pivot]
	s[pivot] = s[counter]
	s[counter] = temp

	return counter
}

func quicksort(s []int, begin int, end int) {
	if end <= begin {
		return
	}

	pivot := partition(s, begin, end)
	quicksort(s, begin, pivot-1)
	quicksort(s, pivot+1, end)
}

func BenchmarkSortIntsV1(b *testing.B) {
	testCases := []int{10, 100, 1000, 10000, 100000}
	for _, tc := range testCases {
		b.Run(fmt.Sprintf("%d", tc), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Stop benchmark timer while preparing the input
				b.StopTimer()
				s := gen(tc)
				b.StartTimer()

				SortIntsV1(s)
			}
		})
	}
}

// Sort slice using Golang sort implementation
func SortIntsV2(s []int) {
	sort.Ints(s)
}

func BenchmarkSortIntsV2(b *testing.B) {
	testCases := []int{10, 100, 1000, 10000, 100000}
	for _, tc := range testCases {
		b.Run(fmt.Sprintf("%d", tc), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Stop benchmark timer while preparing the input
				b.StopTimer()
				s := gen(tc)
				b.StartTimer()

				SortIntsV2(s)
			}
		})
	}
}
