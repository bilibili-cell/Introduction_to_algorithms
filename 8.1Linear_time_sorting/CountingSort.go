package Linear_time_sorting

func CountingSort(a *[]int, k int) {
	listArray := *a
	listArrayCopy := make([]int, len(*a))
	assistArray := make([]int, k+1)
	copy(listArrayCopy, listArray)
	for i := 0; i < len(listArray); i++ {
		assistArray[listArray[i]]++
	}
	for i := 1; i <= k; i++ {
		assistArray[i] += assistArray[i-1]
	}
	for i := len(listArray) - 1; i >= 0; i-- {
		listArray[assistArray[listArrayCopy[i]]-1] = listArrayCopy[i]
		assistArray[listArrayCopy[i]]--
	}
}
