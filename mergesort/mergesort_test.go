package main

import (
	"math/rand"
	"testing"
)

func TestMerge(t *testing.T) {
	tables := []struct {
		l1 []int
		l2 []int
		a  []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{3}, []int{2}, []int{2, 3}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{5, 7, 10}, []int{2, 6, 8}, []int{2, 5, 6, 7, 8, 10}},
	}

	for _, table := range tables {
		ans := merge(table.l1, table.l2)
		if !intSlicesEqual(ans, table.a) {
			t.Errorf("merge of slices %v and %v incorrect, returned %v, expected %v", table.l1, table.l2, ans, table.a)
		}
	}
}

func TestConMergeSort(t *testing.T) {
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
		ans := ConMergeSort(table.q)
		if !intSlicesEqual(ans, table.a) {
			t.Errorf("Sort of slice %v incorrect, returned %v, expected %v", table.q, ans, table.a)
		}
	}
}

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

func BenchmarkMergeSort10(b *testing.B)        { benchmarkMergeSort(10, b) }
func BenchmarkMergeSort100(b *testing.B)       { benchmarkMergeSort(100, b) }
func BenchmarkMergeSort1000(b *testing.B)      { benchmarkMergeSort(1000, b) }
func BenchmarkMergeSort10000(b *testing.B)     { benchmarkMergeSort(10000, b) }
func BenchmarkMergeSort100000(b *testing.B)    { benchmarkMergeSort(100000, b) }
func BenchmarkMergeSort1000000(b *testing.B)   { benchmarkMergeSort(1000000, b) }
func BenchmarkMergeSort100000000(b *testing.B) { benchmarkMergeSort(100000000, b) }

func benchmarkConMergeSort(s int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		testSlice := []int{}
		for i := 0; i < s; i++ {
			testSlice = append(testSlice, rand.Int())
		}
		ConMergeSort(testSlice)
	}
}

func BenchmarkConMergeSort10(b *testing.B)        { benchmarkMergeSort(10, b) }
func BenchmarkConMergeSort100(b *testing.B)       { benchmarkMergeSort(100, b) }
func BenchmarkConMergeSort1000(b *testing.B)      { benchmarkMergeSort(1000, b) }
func BenchmarkConMergeSort10000(b *testing.B)     { benchmarkMergeSort(10000, b) }
func BenchmarkConMergeSort100000(b *testing.B)    { benchmarkMergeSort(100000, b) }
func BenchmarkConMergeSort1000000(b *testing.B)   { benchmarkMergeSort(1000000, b) }
func BenchmarkConMergeSort100000000(b *testing.B) { benchmarkMergeSort(100000000, b) }

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
