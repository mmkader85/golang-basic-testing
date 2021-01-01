package services

import "testing"

// Integration test.
// Sort func in services calling a func in a different layer i.e., utils
func TestSort(t *testing.T) {
	elements := []int{1, 3, 5, 7, 9, 0, 2, 4, 6, 8}
	Sort(elements)

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("last element should be 9")
	}
}
