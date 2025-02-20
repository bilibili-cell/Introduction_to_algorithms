package Hash_Table

import "fmt"

type DirectAddressTable struct {
	table []*int
	size  int
}

func NewDirectAddressTable(size int) *DirectAddressTable {
	return &DirectAddressTable{
		table: make([]*int, size),
		size:  size,
	}
}

func (d *DirectAddressTable) Insert(key int, value int) error {
	if key < 0 || key >= d.size {
		return fmt.Errorf("key %d out of range", key)
	}
	d.table[key] = &value
	return nil
}

func (d *DirectAddressTable) Search(key int) (*int, error) {
	if key < 0 || key >= d.size {
		return nil, fmt.Errorf("key %d out of range", key)
	}
	return d.table[key], nil
}

func (d *DirectAddressTable) Delete(key int) error {
	if key < 0 || key >= d.size {
		return fmt.Errorf("key %d out of range", key)
	}
	d.table[key] = nil
	return nil
}
