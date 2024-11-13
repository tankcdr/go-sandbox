package hash

import (
	"fmt"
)

const NONE = -1

const SLOT_USED = true
const SLOT_EMPTY = false

const LINEAR_PROBE = 1

type LinearProbingHashTableStore[T any] struct {
	key     string
	value   T
	used    bool
	deleted bool
	next    int
}

type LinearProbingHashTable[T any] struct {
	capacity int
	buckets  []*LinearProbingHashTableStore[T]
}

func NewLinearProbingHashTable[T any](capacity int) *LinearProbingHashTable[T] {
	buckets := make([]*LinearProbingHashTableStore[T], capacity)
	for i := range buckets {
		buckets[i] = &LinearProbingHashTableStore[T]{key: EMPTY, next: NONE, used: SLOT_EMPTY}
	}
	return &LinearProbingHashTable[T]{capacity, buckets}
}

func (lpht *LinearProbingHashTable[T]) Capacity() int {
	return lpht.capacity
}

// Set adds a key-value pair to the hash table.
func (lpht *LinearProbingHashTable[T]) Set(key string, value T) {
	index, prev, _ := lpht.findSlot(key)

	if index == NONE {
		panic("Hash table is full")
	}

	if lpht.buckets[index].key == EMPTY {

		lpht.buckets[index].next = NONE
		lpht.buckets[index].key = key
		lpht.buckets[index].value = value
		lpht.buckets[index].used = SLOT_USED
		lpht.buckets[index].deleted = false

		if prev != NONE {
			lpht.buckets[prev].next = index
		}
	}

	lpht.buckets[index].value = value

}

// Get retrieves a value from the hash table.
func (lpht *LinearProbingHashTable[T]) Get(key string) T {
	var value T

	index, _, _ := lpht.findSlot(key)

	if lpht.buckets[index].key == key {
		value = lpht.buckets[index].value
	}

	return value
}

// Delete deletes a key-value pair from the hash table.
func (lpht *LinearProbingHashTable[T]) Delete(key string) bool {
	index, prev, _ := lpht.findSlot(key)
	var emptyValue T

	if lpht.buckets[index].key == key {

		if prev != NONE {
			lpht.buckets[prev].next = lpht.buckets[index].next
		}

		lpht.buckets[index].key = EMPTY
		lpht.buckets[index].value = emptyValue
		lpht.buckets[index].used = SLOT_EMPTY
		lpht.buckets[index].deleted = true

		return true
	}

	return false
}

// Contains checks if a key exists in the hash table.
func (lpht *LinearProbingHashTable[T]) Contains(key string) bool {
	index, _, _ := lpht.findSlot(key)

	if lpht.buckets[index].key == key {
		return true
	}

	return false
}

// Find returns the index of the bucket and the index of the key-value pair in the bucket.
func (lpht *LinearProbingHashTable[T]) Find(key string) (int, int) {
	index, _, probeLength := lpht.findSlot(key)

	if lpht.buckets[index].key != key {
		return NONE, probeLength
	}

	return index, probeLength
}

// Clear removes all key-value pairs from the hash table.
func (lpht *LinearProbingHashTable[T]) Clear() {
	var value T
	for _, bucket := range lpht.buckets {
		bucket.key = EMPTY
		bucket.value = value
		bucket.next = NONE
		bucket.used = SLOT_EMPTY
	}
}

// Dump prints the contents of the hash table.
func (lpht *LinearProbingHashTable[T]) Dump() {
	for index, bucket := range lpht.buckets {
		fmt.Printf("Bucket %d: %v\n", index, bucket)
	}
}

func (lpht *LinearProbingHashTable[T]) findSlot(key string) (slot int, previous int, probeLength int) {
	index := Hash_djb2(key) % lpht.capacity
	prev := NONE
	length := 0
	guard := false

	//case not empty
	for {
		// case we have reached the end of the table
		// break and return not found
		// TODO: wrap around once
		if index >= lpht.capacity {
			if guard {
				panic("Hash table is full")
			}
			index = 0
			guard = true
			continue
		}

		if lpht.buckets[index].used == SLOT_EMPTY && lpht.buckets[index].next != NONE {
			index = lpht.buckets[index].next
			prev = index
			break
		}

		if lpht.buckets[index].used == SLOT_EMPTY {
			break
		}

		//if current key is the key we are looking for
		if lpht.buckets[index].used != SLOT_EMPTY && lpht.buckets[index].key == key {

			break
		}

		prev = index

		if lpht.buckets[index].next != NONE {
			index = lpht.buckets[index].next
		} else {
			index += LINEAR_PROBE
		}
		length++

	}

	//case empty or solution found
	return index, prev, length
}

// Make a display showing whether each array entry is nil.
func (lpht *LinearProbingHashTable[T]) DumpConcise() {
	// Loop through the array.
	for i, bucket := range lpht.buckets {
		if bucket.deleted {
			fmt.Printf("X")
		} else if bucket.used == SLOT_EMPTY {
			// This spot is empty.
			fmt.Printf(".")
		} else {
			// Display this entry.
			fmt.Printf("O")
		}
		if i%50 == 49 {
			fmt.Println()
		}
	}
	fmt.Println()
}

// Return the average probe sequence length for the items in the table.
func (lpht *LinearProbingHashTable[T]) AveProbeSequenceLength() float32 {
	totalLength := 0
	numValues := 0
	for _, bucket := range lpht.buckets {
		if bucket.used {
			_, _, probeLength := lpht.findSlot(bucket.key)
			totalLength += probeLength
			numValues++
		}
	}
	return float32(totalLength) / float32(numValues)
}

func (lpht *LinearProbingHashTable[T]) Probe(key string) int {
	// Hash the key.
	hash := Hash_djb2(key) % lpht.capacity
	fmt.Printf("Probing %s (%d)\n", key, hash)

	// Keep track of a deleted spot if we find one.
	deletedIndex := -1

	// Probe up to lpht.capacity times.
	for i := 0; i < lpht.capacity; i++ {
		index := (hash + i) % lpht.capacity

		fmt.Printf("    %d: ", index)
		if lpht.buckets[index] == nil {
			fmt.Printf("---\n")
		} else if lpht.buckets[index].deleted {
			fmt.Printf("xxx\n")
		} else {
			fmt.Printf("%v\n", lpht.buckets[index])
		}

		// If this spot is empty, the value isn't in the table.
		if lpht.buckets[index] == nil {
			// If we found a deleted spot, return its index.
			if deletedIndex >= 0 {
				fmt.Printf("    Returning deleted index %d\n", deletedIndex)
				return deletedIndex
			}

			// Return this index, which holds nil.
			fmt.Printf("    Returning nil index %d\n", index)
			return index
		}

		// If this spot is deleted, remember where it is.
		if lpht.buckets[index].deleted {
			if deletedIndex < 0 {
				deletedIndex = index
			}
		} else if lpht.buckets[index].key == key {
			// If this cell holds the key, return its data.
			fmt.Printf("    Returning found index %d\n", index)
			return index
		}

		// Otherwise continue the loop.
	}

	// If we get here, then the key is not
	// in the table and the table is full.

	// If we found a deleted spot, return it.
	if deletedIndex >= 0 {
		fmt.Printf("    Returning deleted index %d\n", deletedIndex)
		return deletedIndex
	}

	// There's nowhere to put a new entry.
	fmt.Printf("    Table is full\n")
	return -1
}
