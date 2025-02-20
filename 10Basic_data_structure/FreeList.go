package Basic_data_structure

import (
	"errors"
)

type FreeBlock struct {
	Size int
	Next *FreeBlock
}

type FreeList struct {
	Head *FreeBlock
}

func NewFreeList() *FreeList {
	return &FreeList{Head: nil}
}

func (f *FreeList) Alloc(size int) (*FreeBlock, error) {
	prev := &f.Head
	for block := f.Head; block != nil; block = block.Next {
		if block.Size >= size {
			if block.Size > size {
				newBlock := &FreeBlock{
					Size: block.Size - size,
					Next: block.Next,
				}
				block.Size = size
				block.Next = newBlock
			}
			*prev = block.Next
			return block, nil
		}
		prev = &block.Next
	}
	return nil, errors.New("out of memory")
}
func (f *FreeList) Free(block *FreeBlock) {
	block.Next = f.Head
	f.Head = block
}
