package sort

import (
	"log"
	"sort"
)

// Sort elements in ascending order
func BubbleSort(elements []int) {
	log.Println("Sorting using BubbleSort")
	keepWorking := true
	for keepWorking {
		keepWorking = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				keepWorking = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}

// Golang built-in sort function
func BuiltInSort(elements []int) {
	log.Println("Sorting using BuiltInSort")
	sort.Ints(elements)
}

func GetIntElements(n int) []int {
	results := make([]int, n)

	j := 0
	for i := n; i > 1; i-- {
		results[j] = i
		j++
	}

	return results
}
