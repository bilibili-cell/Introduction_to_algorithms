package HeapSort

import (
	"math"
)

func Parent(i int) int {
	return int(math.Floor(float64(i) / 2))
}

func Left(i int) int {
	return 2 * i
}

func Right(i int) int {
	return 2*i + 1
}

func MaxHeapify(a *[]int, i int) {
	listArray := *a
	l := Left(i)
	r := Right(i)
	largest := i
	if l <= len(listArray) && listArray[l] > listArray[largest] {
		largest = l
	}
	if r <= len(listArray) && listArray[r] > listArray[largest] {
		largest = r
	}
	if largest != i {
		listArray[largest], listArray[i] = listArray[i], listArray[largest]
		MaxHeapify(a, largest)
	}
}

func BuildMaxHeap(a *[]int) {
	listArray := *a
	if len(listArray) < 2 {
		return
	} else {
		for i := len(listArray)/2 - 1; i >= 0; i-- {
			MaxHeapify(a, i)
		}
	}
}
