package Median_and_order_statistics

func Minimum(arr *[]int) int {
	listArray := *arr
	min := listArray[0]
	for i := 1; i < len(listArray); i++ {
		if listArray[i] < min {
			min = listArray[i]
		}
	}
	return min

}