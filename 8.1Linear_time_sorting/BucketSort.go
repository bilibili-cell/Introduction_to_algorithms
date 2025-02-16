package Linear_time_sorting

type Node struct {
	value int
	Next  *Node
}

func BucketSort(arr *[]int) {
	listArray := *arr
	assistArray := make([]Node, 10)
	for i := 0; i < 10; i++ {
		assistArray[i] = Node{
			value: 0,
			Next:  nil,
		}
	}
	for _, item := range listArray {
		node := Node{
			value: item,
			Next:  nil,
		}
		value := assistArray[item/10].value
		listNode := assistArray[item/10].Next
		pre := &assistArray[item/10]
		for i := 0; i < value; i++ {
			if listNode != nil {
				v := listNode.value
				if node.value < v {
					node.Next = pre.Next
					pre.Next = &node
				} else {
					pre = listNode
					listNode = listNode.Next
				}
			} else {
				break
			}
		}
		assistArray[item/10].value++

	}
	i := 0
	for _, item := range assistArray {

		if item.Next != nil {
			for j := 1; j <= item.value; j++ {
				listArray[i] = item.Next.value
				item = *item.Next
				i++
			}
		}

	}

}
