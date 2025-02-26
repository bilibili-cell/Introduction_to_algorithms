package Hash_Table

const TableSize = 10

type HashTable struct {
	table [TableSize]int
}

func (ht *HashTable) Insert(key int) {
	index := key % TableSize
	for ht.table[index] != 0 {
		index = (index + 1) % TableSize
	}
	ht.table[index] = key
}

func (ht *HashTable) Search(key int) int {
	index := key % TableSize
	for ht.table[index] != 0 {
		if ht.table[index] == key {
			return index
		}
		index = (index + 1) % TableSize
	}
	return -1
}

func (ht *HashTable) Delete(key int) {
	index := key % TableSize
	for ht.table[index] != 0 {
		if ht.table[index] == key {
			ht.table[index] = 0
			return
		}
		index = (index + 1) % TableSize
	}
}
