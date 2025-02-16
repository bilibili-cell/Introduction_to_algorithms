package __1Insertion_sort

// 当 key 小于所有已排序元素时（即需要插入到数组开头），你的代码并没有在内层循环结束后将 key 赋值给 arr[0]。这样会导致 key 丢失，最终数组的第一个元素被重复的值覆盖。
func InsertionSort(arr []int, ascending bool) []int {
	length := len(arr)
	if length < 2 {
		return arr
	} else if ascending {
		for i := 1; i < length; i++ {
			key := arr[i]
			for j := i - 1; j >= 0; j-- {
				if key < arr[j] {
					arr[j+1] = arr[j]
					arr[j] = key
				} else {
					arr[j+1] = key
					break
				}

			}
		}
	} else if !ascending {
		for i := 1; i < length; i++ {
			key := arr[i]
			for j := i - 1; j >= 0; j-- {
				if key > arr[j] {
					arr[j+1] = arr[j]
					arr[j] = key
				} else {
					arr[j+1] = key
					break
				}

			}
		}
	}
	return arr
}
func InsertionSortByAI(arr []int, ascending bool) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}

	for i := 1; i < length; i++ {
		key := arr[i]
		j := i - 1

		if ascending {
			for j >= 0 && arr[j] > key {
				arr[j+1] = arr[j]
				j--
			}
		} else {
			for j >= 0 && arr[j] < key {
				arr[j+1] = arr[j]
				j--
			}
		}
		arr[j+1] = key
	}
	return arr
}
