package main

import (
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tables := []struct {
		q []int
		a []int
	}{
		{[]int{}, []int{}},
		{[]int{3}, []int{3}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{2, 1, 3}, []int{1, 2, 3}},
		{[]int{3, 1, 2}, []int{1, 2, 3}},
		{[]int{5, 7, 10, 22, 1, 6}, []int{1, 5, 6, 7, 10, 22}},
		{[]int{1, 5, 7, 10, 22, 1, 6}, []int{1, 1, 5, 6, 7, 10, 22}},
		{[]int{3, 5, 7, 10, 22, 1, 6}, []int{1, 3, 5, 6, 7, 10, 22}},
	}

	for _, table := range tables {
		ans := MergeSort(table.q)
		if !intSlicesEqual(ans, table.a) {
			t.Errorf("Sort of slice %v incorrect, returned %v, expected %v", table.q, ans, table.a)
		}
	}
}

func benchmarkMergeSort(s int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		testSlice := []int{}
		for i := 0; i < s; i++ {
			testSlice = append(testSlice, rand.Int())
		}
		MergeSort(testSlice)
	}
}

func BenchmarkMergeSort10(b *testing.B)    { benchmarkMergeSort(10, b) }
func BenchmarkMergeSort100(b *testing.B)   { benchmarkMergeSort(100, b) }
func BenchmarkMergeSort1000(b *testing.B)  { benchmarkMergeSort(1000, b) }
func BenchmarkMergeSort10000(b *testing.B) { benchmarkMergeSort(10000, b) }

func intSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
