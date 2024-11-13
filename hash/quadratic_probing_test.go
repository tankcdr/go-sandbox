package hash_test

import (
	"fmt"
	"testing"

	"github.com/tankcdr/hash"
)

func TestQuadraticProbingHashTable_Create(t *testing.T) {
	t.Parallel()

	qpht := hash.NewQuadraticProbingHashTable[int](100)
	if qpht == nil {
		t.Errorf("Failed to create chaining hash table")
	}

	if qpht.Capacity() != 100 {
		t.Errorf("Expected 100 buckets, got %d", qpht.Capacity())
	}
}

func TestQuadraticProbingHashTable_Set(t *testing.T) {
	t.Parallel()

	qpht := hash.NewQuadraticProbingHashTable[int](10)

	qpht.Set("one", 1)
	qpht.Set("two", 2)
	qpht.Set("three", 3)

	if bucket_index, _ := qpht.Find("one"); bucket_index == -1 {
		t.Errorf("Failed to find 1 in the hash table")
	}

	if bucket_index, _ := qpht.Find("two"); bucket_index == -1 {
		t.Errorf("Failed to find 2 in the hash table")
	}

	if bucket_index, _ := qpht.Find("three"); bucket_index == -1 {
		t.Errorf("Failed to find 3 in the hash table")
	}
}

func TestQuadraticProbingHashTable_Remove(t *testing.T) {
	t.Parallel()

	qpht := hash.NewQuadraticProbingHashTable[int](10)

	qpht.Set("one", 1)
	qpht.Set("two", 2)
	qpht.Set("three", 3)

	qpht.Delete("one")

	if bucket_index, _ := qpht.Find("one"); bucket_index != -1 {
		t.Errorf("Failed to find 1 in the hash table")
	}

	if bucket_index, _ := qpht.Find("two"); bucket_index == -1 {
		t.Errorf("Failed to find 2 in the hash table")
	}

	if bucket_index, _ := qpht.Find("three"); bucket_index == -1 {
		t.Errorf("Failed to find 3 in the hash table")
	}

}

func TestQuadraticProbingHashTable_Clear(t *testing.T) {
	t.Parallel()

	qpht := hash.NewQuadraticProbingHashTable[int](10)

	qpht.Set("one", 1)
	qpht.Set("two", 2)
	qpht.Set("three", 3)

	qpht.Clear()

	if bucket_index, _ := qpht.Find("one"); bucket_index != -1 {
		fmt.Printf("bucket_index: %d\n", bucket_index)
		t.Errorf("Failed to find 1 in the hash table")
	}

	if bucket_index, _ := qpht.Find("two"); bucket_index != -1 {
		t.Errorf("Failed to find 2 in the hash table")
	}

	if bucket_index, _ := qpht.Find("three"); bucket_index != -1 {
		t.Errorf("Failed to find 3 in the hash table")
	}
}

func TestQuadraticProbingHashTable_Find(t *testing.T) {
	t.Parallel()

	qpht := hash.NewQuadraticProbingHashTable[int](10)

	qpht.Set("one", 1)
	qpht.Set("two", 2)
	qpht.Set("three", 3)

	if bucket_index, _ := qpht.Find("one"); bucket_index == -1 {
		t.Errorf("Failed to find 1 in the hash table")
	}

	if bucket_index, _ := qpht.Find("two"); bucket_index == -1 {
		t.Errorf("Failed to find 2 in the hash table")
	}

	if bucket_index, _ := qpht.Find("three"); bucket_index == -1 {
		t.Errorf("Failed to find 3 in the hash table")
	}
}

func TestQuadraticProbingHashTable_FindWithStruct(t *testing.T) {
	t.Parallel()

	type Employee struct {
		Name  string
		Phone string
	}

	qpht := hash.NewQuadraticProbingHashTable[Employee](10)

	a := Employee{"John Doe", "555-1234"}
	b := Employee{"Jane Doe", "555-6789"}
	c := Employee{"John Smith", "555-0000"}

	qpht.Set(a.Name, a)
	qpht.Set(b.Name, b)
	qpht.Set(c.Name, c)

	if bucket_index, _ := qpht.Find(a.Name); bucket_index == -1 {
		t.Errorf("Failed to find 1 in the hash table")
	}

	if bucket_index, _ := qpht.Find(b.Name); bucket_index == -1 {
		t.Errorf("Failed to find 2 in the hash table")
	}

	if bucket_index, _ := qpht.Find(c.Name); bucket_index == -1 {
		t.Errorf("Failed to find 3 in the hash table")
	}
}

func TestQuadraticProbingHashTable_Get(t *testing.T) {
	t.Parallel()

	qpht := hash.NewQuadraticProbingHashTable[int](10)

	qpht.Set("one", 1)
	qpht.Set("two", 2)
	qpht.Set("three", 3)

	if qpht.Get("one") != 1 {
		t.Errorf("Failed to get 1 from the hash table")
	}

	if qpht.Get("two") != 2 {
		t.Errorf("Failed to get 2 from the hash table")
	}

	if qpht.Get("three") != 3 {
		t.Errorf("Failed to get 3 from the hash table")
	}
}

func TestQuadraticProbingHashTable_Contains(t *testing.T) {
	t.Parallel()

	qpht := hash.NewQuadraticProbingHashTable[int](10)

	qpht.Set("one", 1)
	qpht.Set("two", 2)
	qpht.Set("three", 3)

	if !qpht.Contains("one") {
		t.Errorf("Failed to find 1 in the hash table")
	}

	if !qpht.Contains("two") {
		t.Errorf("Failed to find 2 in the hash table")
	}

	if !qpht.Contains("three") {
		t.Errorf("Failed to find 3 in the hash table")
	}
}
