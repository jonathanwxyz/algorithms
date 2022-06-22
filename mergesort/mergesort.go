package main

import (
	"fmt"
	"sync"
)

type Comparable interface {
	string | int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uintptr
}

func ConMergeSort[C Comparable](list []C) []C {
	if len(list) > 1 {
		var half int
		if len(list)%2 == 1 {
			half = (len(list) + 1) / 2
		} else {
			half = len(list) / 2
		}
		var l1, l2 []C
		var waitGroup sync.WaitGroup
		waitGroup.Add(2)
		go func() {
			defer waitGroup.Done()
			l1 = ConMergeSort(list[0:half])
		}()
		go func() {
			defer waitGroup.Done()
			l2 = ConMergeSort(list[half:])
		}()
		waitGroup.Wait()
		return merge(l1, l2)
	}
	return list
}

func MergeSort[C Comparable](list []C) []C {
	if len(list) > 1 {
		var half int
		if len(list)%2 == 1 {
			half = (len(list) + 1) / 2
		} else {
			half = len(list) / 2
		}
		return merge(MergeSort(list[0:half]), MergeSort(list[half:]))
	}
	return list
}

func merge[C Comparable](l1, l2 []C) []C {
	sorted := []C{}
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
