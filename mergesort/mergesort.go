package main

/*
type Number interface {
    int64 | float64
}
*/

func MergeSort(list []int) []int {
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

func merge(l1, l2 []int) []int {
	sorted := []int{}
	i, j := 0, 0
	for {
		if i == len(l1) {
			for ; j < len(l2); j++ {
				sorted = append(sorted, l2[j])
			}
		} else if j == len(l2) {
			for ; i < len(l1); i++ {
				sorted = append(sorted, l1[i])
			}
		} else {
			if l1[i] < l2[j] {
				sorted = append(sorted, l1[i])
				i++
			} else {
				sorted = append(sorted, l2[j])
				j++
			}
		}
		if i == len(l1) && j == len(l2) {
			break
		}
	}
	return sorted
}
