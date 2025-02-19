package Basic_data_structure

import (
	"errors"
)

type Node struct {
	Prev, Next *Node
	Key        int
}

type DoubleLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (d *DoubleLinkedList) DoubleListDelete(pos int) error {
	if pos < 0 || pos >= d.Length {
		return errors.New("position out of range")
	}

	if pos == 0 {
		if d.Head == d.Tail {
			d.Head = nil
			d.Tail = nil
		} else {
			d.Head = d.Head.Next
			d.Head.Prev = nil
		}
	} else if pos == d.Length-1 {
		if d.Head == d.Tail {
			d.Head = nil
			d.Tail = nil
		} else {
			d.Tail = d.Tail.Prev
			d.Tail.Next = nil
		}
	} else {
		curNode := d.Head
		for i := 0; i < pos; i++ {
			curNode = curNode.Next
		}
		curNode.Prev.Next = curNode.Next
		curNode.Next.Prev = curNode.Prev
	}

	d.Length--
	return nil
}

func (d *DoubleLinkedList) DoubleListInsert(param ...int) error {
	if len(param) == 0 {
		return errors.New("invalid parameters")
	}

	node := &Node{
		Key:  param[0],
		Prev: nil,
		Next: nil,
	}

	if len(param) == 1 {
		if d.Tail == nil {
			d.Head = node
			d.Tail = node
		} else {
			d.Tail.Next = node
			node.Prev = d.Tail
			d.Tail = node
		}
	} else if len(param) == 2 {
		pos := param[1]
		if pos < 0 || pos > d.Length {
			return errors.New("position out of range")
		}

		if pos == 0 {
			if d.Head == nil {
				d.Head = node
				d.Tail = node
			} else {
				node.Next = d.Head
				d.Head.Prev = node
				d.Head = node
			}
		} else if pos == d.Length {
			d.Tail.Next = node
			node.Prev = d.Tail
			d.Tail = node
		} else {
			curNode := d.Head
			for i := 0; i < pos; i++ {
				curNode = curNode.Next
			}
			node.Prev = curNode.Prev
			node.Next = curNode
			curNode.Prev.Next = node
			curNode.Prev = node
		}
	} else {
		return errors.New("invalid parameters")
	}

	d.Length++
	return nil
}
func (d *DoubleLinkedList) DoubleLinkedListSearch(pos int) (int, error) {
	if pos < 0 || pos >= d.Length {
		return 0, errors.New("position out of range")
	}

	curNode := d.Head
	for i := 0; i < pos; i++ {
		curNode = curNode.Next
	}

	return curNode.Key, nil
}
