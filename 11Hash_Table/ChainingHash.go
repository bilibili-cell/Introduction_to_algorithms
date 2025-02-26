package Hash_Table

import (
	"errors"
)

type Node struct {
	Value int
	Next  *Node
}

type ChainingHash struct {
	List []*Node
	Size int
}

func NewChainingHash() *ChainingHash {
	return &ChainingHash{
		List: make([]*Node, 10),
		Size: 10,
	}
}

func Hash(i int) int {
	return i % 10
}

func (chainingHash *ChainingHash) Insert(i int) {
	node := &Node{
		Value: i,
		Next:  nil,
	}
	headNode := chainingHash.List[i%10]
	node.Next = headNode
	chainingHash.List[i%10] = node
}

func (chainingHash *ChainingHash) Search(i int) (int, error) {
	headNode := chainingHash.List[i%10]
	for headNode != nil {
		if headNode.Value == i {
			return headNode.Value, nil
		}
		headNode = headNode.Next
	}
	return 0, errors.New("value not found")
}

func (chainingHash *ChainingHash) Remove(i int) error {
	headNode := chainingHash.List[i%10]
	if headNode == nil {
		return errors.New("value not found")
	}
	if headNode.Value == i {
		chainingHash.List[i%10] = headNode.Next
		return nil
	}
	prevNode := headNode
	for headNode != nil {
		if headNode.Value == i {
			prevNode.Next = headNode.Next
			return nil
		}
		prevNode = headNode
		headNode = headNode.Next
	}
	return errors.New("value not found")
}
