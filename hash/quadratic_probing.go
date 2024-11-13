package hash

import (
	"fmt"
)

const GUARD = 5

type QuadraticProbingHashTableStore[T any] struct {
	key     string
	value   T
	deleted bool
}

type QuadraticProbingHashTable[T any] struct {
	capacity int
	buckets  []*QuadraticProbingHashTableStore[T]
}

func NewQuadraticProbingHashTable[T any](capacity int) *QuadraticProbingHashTable[T] {
	buckets := make([]*QuadraticProbingHashTableStore[T], capacity)
	for i := range buckets {
		buckets[i] = &QuadraticProbingHashTableStore[T]{key: EMPTY, deleted: NOT_DELETED}
	}
	return &QuadraticProbingHashTable[T]{capacity, buckets}
}

func (qpht *QuadraticProbingHashTable[T]) Capacity() int {
	return qpht.capacity
}

// Set adds a key-value pair to the hash table.
func (qpht *QuadraticProbingHashTable[T]) Set(key string, value T) {
	index, _ := qpht.findSlot(key)

	if index == NONE {
		panic("Hash table is full")
	}

	if qpht.buckets[index].key == EMPTY {
		qpht.buckets[index].key = key
		qpht.buckets[index].deleted = USED
	}

	qpht.buckets[index].value = value

}

// Get retrieves a value from the hash table.
func (qpht *QuadraticProbingHashTable[T]) Get(key string) T {
	var value T

	index, _ := qpht.findSlot(key)

	if qpht.buckets[index].key == key {
		value = qpht.buckets[index].value
	}

	return value
}

// Delete deletes a key-value pair from the hash table.
func (qpht *QuadraticProbingHashTable[T]) Delete(key string) bool {
	var emptyValue T

	index, _ := qpht.findSlot(key)

	if qpht.buckets[index].key == key {
		qpht.buckets[index].key = EMPTY
		qpht.buckets[index].value = emptyValue
		qpht.buckets[index].deleted = DELETED

		return true
	}

	return false
}

// Contains checks if a key exists in the hash table.
func (qpht *QuadraticProbingHashTable[T]) Contains(key string) bool {
	index, _ := qpht.findSlot(key)

	if index != NONE && qpht.buckets[index].key == key {
		return true
	}

	return false
}

// Find returns the index of the bucket and the index of the key-value pair in the bucket.
func (qpht *QuadraticProbingHashTable[T]) Find(key string) (int, int) {
	index, probeLength := qpht.findSlot(key)

	if index != NONE && qpht.buckets[index].key != key {
		return NONE, probeLength
	}

	return index, probeLength
}

// Clear removes all key-value pairs from the hash table.
func (qpht *QuadraticProbingHashTable[T]) Clear() {
	var value T

	for _, bucket := range qpht.buckets {
		bucket.key = EMPTY
		bucket.value = value
		bucket.deleted = NOT_DELETED
	}
}

// Dump prints the contents of the hash table.
func (qpht *QuadraticProbingHashTable[T]) Dump() {
	for index, bucket := range qpht.buckets {
		fmt.Printf("Bucket %d: %v\n", index, bucket)
	}
}

func (qpht *QuadraticProbingHashTable[T]) findSlot(key string) (slot int, probeLength int) {
	slot = NONE
	probeLength = 0
	guard := false

	//case not empty
	for i := 0; ; i++ {
		//quadradic probing
		index := (Hash_djb2(key) + i*i) % qpht.capacity

		// case we have reached the end of the table
		// wrap around once
		if i >= qpht.capacity {
			if guard {
				break
			}
			guard = true
			continue
		}

		probeLength = i

		if qpht.buckets[index].deleted == DELETED {
			continue
		}

		if qpht.buckets[index].key == key || qpht.buckets[index].key == EMPTY {
			slot = index
			break
		}
	}

	//case empty or solution found
	return slot, probeLength
}

// Make a display showing whether each array entry is nil.
func (qpht *QuadraticProbingHashTable[T]) DumpConcise() {
	// Loop through the array.
	for i, bucket := range qpht.buckets {
		if bucket.deleted {
			fmt.Printf("X")
		} else if bucket.key != EMPTY {
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
func (qpht *QuadraticProbingHashTable[T]) AveProbeSequenceLength() float32 {
	totalLength := 0
	numValues := 0
	for _, bucket := range qpht.buckets {
		if bucket.key != EMPTY {
			_, probeLength := qpht.findSlot(bucket.key)
			totalLength += probeLength
			numValues++
		}
	}
	return float32(totalLength) / float32(numValues)
}

func (qpht *QuadraticProbingHashTable[T]) Probe(key string) int {
	// Hash the key.
	hash := Hash_djb2(key) % qpht.capacity
	fmt.Printf("Probing %s (%d)\n", key, hash)

	// Keep track of a deleted spot if we find one.
	deletedIndex := -1

	// Probe up to qpht.capacity times.
	for i := 0; i < qpht.capacity; i++ {
		index := (hash + i) % qpht.capacity

		fmt.Printf("    %d: ", index)
		if qpht.buckets[index] == nil {
			fmt.Printf("---\n")
		} else if qpht.buckets[index].deleted {
			fmt.Printf("xxx\n")
		} else {
			fmt.Printf("%v\n", qpht.buckets[index])
		}

		// If this spot is empty, the value isn't in the table.
		if qpht.buckets[index] == nil {
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
		if qpht.buckets[index].deleted {
			if deletedIndex < 0 {
				deletedIndex = index
			}
		} else if qpht.buckets[index].key == key {
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
