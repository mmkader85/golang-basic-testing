package sort

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Unit test for BubbleSort
func TestBubbleSortAscOrder(t *testing.T) {
	elements := []int{1, 3, 5, 7, 9, 0, 2, 4, 6, 8}
	BubbleSort(elements)

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("last element should be 9")
	}
}

// Unit test for BubbleSort
func TestBubbleSortAscOrderLargeElements(t *testing.T) {
	largestNumber := 1000
	elements := GetIntElements(largestNumber)
	BubbleSort(elements)

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != largestNumber {
		t.Errorf("last element should be %d but actual is %d", largestNumber, elements[len(elements)-1])
	}
}

func TestBuiltInSort(t *testing.T) {
	largestNumber := 1000
	elements := GetIntElements(largestNumber)
	BuiltInSort(elements)

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != largestNumber {
		t.Errorf("last element should be %d but actual is %d", largestNumber, elements[len(elements)-1])
	}
}

// It took 563 ns/op for 1,000 elements
// It took 75702 ns/op for 30,000 elements
// It took 1091582654 ns/op for 35,000 elements
// It took 2215449756 ns/op for 50,000 elements
// We can conclude that with more than 30,000, it starts to perform very badly.
// So in our sort_service, we can use that as threshold and choose more suitable function.
func BenchmarkBubbleSort(b *testing.B) {
	largestNumber := 30000
	elements := GetIntElements(largestNumber)

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkBuiltInSort(b *testing.B) {
	// It took 37069 ns/op for 1,000 elements
	// It took 1673564 ns/op for 30,000 elements
	// It took 1932222 ns/op for 35,000 elements
	// It took 2845430 ns/op for 50,000 elements
	largestNumber := 30000
	elements := GetIntElements(largestNumber)

	for i := 0; i < b.N; i++ {
		BuiltInSort(elements)
	}
}

// Unit test for performance of BubbleSort
func TestBubbleSortPerformance(t *testing.T) {
	elements := []int{1, 3, 5, 7, 9, 0, 2, 4, 6, 8}

	timeOutChan := make(chan bool)
	defer close(timeOutChan)

	go func() {
		BubbleSort(elements)
		timeOutChan <- false
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		timeOutChan <- true
	}()

	if <-timeOutChan {
		assert.Fail(t, "BubbleSort took more than 50ms")
		return
	}
}
