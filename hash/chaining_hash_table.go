package hash

import (
	"fmt"

	"github.com/tankcdr/lists"
)

type ChainingHashTableStore[T any] struct {
	key   string
	value T
}

type ChainingHashTable[T any] struct {
	numBuckets int
	buckets    []*lists.DoublyLinkedList[ChainingHashTableStore[T]]
}

func NewChainingHashTable[T any](numBuckets int) *ChainingHashTable[T] {
	buckets := make([]*lists.DoublyLinkedList[ChainingHashTableStore[T]], numBuckets)
	for i := range buckets {
		buckets[i] = lists.NewDoublyLinkedList[ChainingHashTableStore[T]]()
	}
	return &ChainingHashTable[T]{numBuckets: numBuckets, buckets: buckets}
}

func (cht *ChainingHashTable[T]) NumBuckets() int {
	return cht.numBuckets
}

// Set adds a key-value pair to the hash table.
func (cht *ChainingHashTable[T]) Set(key string, value T) {

	found_bucket, found_index := cht.Find(key)

	if found_index == -1 {
		index := createIndex(key, cht.numBuckets)
		cht.buckets[index].Append(ChainingHashTableStore[T]{key, value})
		return
	}

	cht.buckets[found_bucket].RemoveAt(found_index)
	cht.buckets[found_bucket].Append(ChainingHashTableStore[T]{key, value})
}

// Get retrieves a value from the hash table.
func (cht *ChainingHashTable[T]) Get(key string) T {
	var value T

	index := createIndex(key, cht.numBuckets)
	values := cht.buckets[index].Values()

	for _, v := range values {
		if v.key == key {
			value = v.value
			break
		}
	}
	return value
}

// Delete deletes a key-value pair from the hash table.
func (cht *ChainingHashTable[T]) Delete(key string) bool {
	index := createIndex(key, cht.numBuckets)

	values := cht.buckets[index].Values()

	for _, v := range values {
		if v.key == key {
			cht.buckets[index].Remove(v)
			return true
		}
	}
	return false
}

// Contains checks if a key exists in the hash table.
func (cht *ChainingHashTable[T]) Contains(key string) bool {

	if _, index := cht.Find(key); index != -1 {
		return true
	}

	return false
}

// Find returns the index of the bucket and the index of the key-value pair in the bucket.
func (cht *ChainingHashTable[T]) Find(key string) (int, int) {
	bucket_index := -1

	index := createIndex(key, cht.numBuckets)

	if !cht.buckets[index].IsEmpty() {
		values := cht.buckets[index].Values()

		for i, v := range values {
			if key == v.key {
				bucket_index = i
				break
			}
		}
	}

	return index, bucket_index
}

// Clear removes all key-value pairs from the hash table.
func (cht *ChainingHashTable[T]) Clear() {
	for _, bucket := range cht.buckets {
		bucket.Clear()
	}
}

// Dump prints the contents of the hash table.
func (cht *ChainingHashTable[T]) Dump() {
	for index, bucket := range cht.buckets {
		fmt.Printf("Bucket %d:\n", index)
		for _, value := range bucket.Values() {
			fmt.Printf("\t%v\n", value.value)
		}
	}
}

func createIndex(key string, numBuckets int) int {
	return hash(key) % numBuckets
}

// *db2 hash function. See http://www.cse.yorku.ca/~oz/hash.html.
func hash(value string) int {
	hash := 5381
	for _, ch := range value {
		hash = ((hash << 5) + hash) + int(ch)
	}

	// Make sure the result is non-negative.
	if hash < 0 {
		hash = -hash
	}
	return hash
}
