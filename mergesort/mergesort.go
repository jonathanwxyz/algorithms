package main

import (
	"sync"
    "runtime"
)

type Comparable interface {
	string | int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uintptr
}

func MergeSort[C Comparable](arr []C) []C {
    if len(arr) < 512 {
        return SeqMergeSort(arr)
    }
    numCores := runtime.NumCPU()
    return ConLowMergeSort(arr, numCores)
}

// ConLowMergeSort: Concurrent Low-goroutine Merge Sort
func ConLowMergeSort[C Comparable](arr []C, coroutines int) []C {
	if len(arr) <= 1 {
        return arr
    }

    half := len(arr) / 2

    if coroutines <= 0 {
        return merge(SeqMergeSort(arr[0:half]), SeqMergeSort(arr[half:]))
    }

    var l1, l2 []C
    var waitGroup sync.WaitGroup
    waitGroup.Add(2)

    go func() {
        defer waitGroup.Done()
        l1 = ConLowMergeSort(arr[0:half], coroutines - 2)
    }()
    go func() {
        defer waitGroup.Done()
        l2 = ConLowMergeSort(arr[half:], coroutines - 2)
    }()

    waitGroup.Wait()
    return merge(l1, l2)
}

// SeqMergeSort: Sequential Merge Sort
func SeqMergeSort[C Comparable](arr []C) []C {
	if len(arr) > 1 {
		var half int
		if len(arr)%2 == 1 {
			half = (len(arr) + 1) / 2
		} else {
			half = len(arr) / 2
		}
		return merge(SeqMergeSort(arr[0:half]), SeqMergeSort(arr[half:]))
	}
	return arr
}

// ConMergeSort: Concurrent Merge Sort
func ConMergeSort[C Comparable](arr []C) []C {

	if len(arr) <= 1 {
        return arr
    }

    half := len(arr) / 2
    // performance optimisation
    if len(arr) <= 1024 {
        return merge(SeqMergeSort(arr[0:half]), SeqMergeSort(arr[half:]))
    }

    var l1, l2 []C
    var waitGroup sync.WaitGroup
    waitGroup.Add(2)

    go func() {
        defer waitGroup.Done()
        l1 = ConMergeSort(arr[0:half])
    }()
    go func() {
        defer waitGroup.Done()
        l2 = ConMergeSort(arr[half:])
    }()

    waitGroup.Wait()
    return merge(l1, l2)
}

func merge[C Comparable](l1, l2 []C) []C {
    sorted := make([]C, 0, len(l1) + len(l2))
	i, j := 0, 0
	for i != len(l1) || j != len(l2) {
		if i == len(l1) {
			sorted = append(sorted, l2[j:]...)
			break
		} else if j == len(l2) {
			sorted = append(sorted, l1[i:]...)
			break
		} else if l1[i] < l2[j] {
			sorted = append(sorted, l1[i])
			i++
		} else {
			sorted = append(sorted, l2[j])
			j++
		}
	}
	return sorted
}
