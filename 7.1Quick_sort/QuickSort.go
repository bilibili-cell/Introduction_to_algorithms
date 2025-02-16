package QuickSort

func Partition(arr *[]int, left int, right int) int {
	listArray := *arr
	pivot := listArray[right]
	i := left
	for j := left; j < right; j++ {
		if listArray[j] <= pivot {
			listArray[j], listArray[i] = listArray[i], listArray[j]
			i++
		}
	}
	listArray[right], listArray[i] = listArray[i], listArray[right]
	return i
}
func QuickSort(arr *[]int, left int, right int) {
	if left < right {
		partition := Partition(arr, left, right)
		QuickSort(arr, left, partition-1)
		QuickSort(arr, partition+1, right)
	}
}
