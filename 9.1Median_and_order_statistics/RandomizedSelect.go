package Median_and_order_statistics

import (
	"math/rand"
	"time"
)

// 随机化划分函数
func randomizedPartition(a *[]int, low, high int) int {
	arr := *a
	rand.Seed(time.Now().UnixNano())
	pivotIndex := low + rand.Intn(high-low+1)
	arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]

	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// 随机化选择函数
func randomizedSelect(a *[]int, low, high, k int) int {
	arr := *a
	if low == high {
		return arr[low]
	}

	pivotIndex := randomizedPartition(a, low, high)
	length := pivotIndex - low + 1

	if length == k {
		return arr[pivotIndex]
	} else if k < length {
		return randomizedSelect(a, low, pivotIndex-1, k)
	} else {
		return randomizedSelect(a, pivotIndex+1, high, k-length)
	}
}
