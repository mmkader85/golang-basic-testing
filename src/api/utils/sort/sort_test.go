package sort

import "testing"

// Unit test
func TestBubbleSortDescOrder(t *testing.T) {
	elements := []int{1, 3, 5, 7, 9, 0, 2, 4, 6, 8}
	BubbleSort(elements)

	if elements[0] != 9 {
		t.Error("first element should be 9")
	}
	if elements[len(elements)-1] != 0 {
		t.Error("last element should be 0")
	}
}
