package services

import "github.com/mmkader85/golang-basic-testing/src/api/utils/sort"

// Sort the slice of integers using BubbleSort function of utils/sort function
// See
func Sort(elements []int) {
	if len(elements) >= 30000 {
		sort.BuiltInSort(elements)
	} else {
		sort.BubbleSort(elements)
	}
}
