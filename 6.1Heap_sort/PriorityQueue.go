package HeapSort

import (
	"errors"
	"math"
)

type PriorityQueue struct {
	Queue []int
}

func (p PriorityQueue) Maximum() (int, error) {
	if len(p.Queue) == 0 {
		return 0, errors.New("Max heap is empty")
	} else {
		return p.Queue[0], nil
	}
}

func (p PriorityQueue) ExtractMax() (int, error) {
	if len(p.Queue) != 0 {
		max := p.Queue[0]
		p.Queue[0], p.Queue[len(p.Queue)-1] = p.Queue[len(p.Queue)-1], p.Queue[0]
		p.Queue = p.Queue[:len(p.Queue)-1]
		MaxHeapify(&p.Queue, 0)
		return max, nil
	} else {
		return 0, errors.New("heap is empty")
	}
}
func (p PriorityQueue) IncreaseKey(i int, key int) error {
	if key < p.Queue[i] {
		return errors.New("The heap is full")
	} else {
		p.Queue[i] = key
		for i > 0 && p.Queue[Parent(i)] < p.Queue[i] {
			p.Queue[i], p.Queue[Parent(i)] = p.Queue[Parent(i)], p.Queue[i]
			i = Parent(i)
		}
		return nil
	}
}

func (p PriorityQueue) Insert(key int) error {
	p.Queue = append(p.Queue, math.MinInt)
	err := p.IncreaseKey(len(p.Queue)-1, key)
	return err
}
