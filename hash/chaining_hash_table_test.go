package hash_test

import (
	"testing"

	"github.com/tankcdr/hash"
)

func TestChainingHashTable_Create(t *testing.T) {
	t.Parallel()

	cht := hash.NewChainingHashTable[int](10)
	if cht == nil {
		t.Errorf("Failed to create chaining hash table")
	}

	if cht.NumBuckets() != 10 {
		t.Errorf("Expected 10 buckets, got %d", cht.NumBuckets())
	}
}

func TestChainingHashTable_Set(t *testing.T) {
	t.Parallel()

	cht := hash.NewChainingHashTable[int](10)

	cht.Set("one", 1)
	cht.Set("two", 2)
	cht.Set("three", 3)

	if _, bucket_index := cht.Find("one"); bucket_index == -1 {
		t.Errorf("Failed to find 1 in the hash table")
	}

	if _, bucket_index := cht.Find("two"); bucket_index == -1 {
		t.Errorf("Failed to find 2 in the hash table")
	}

	if _, bucket_index := cht.Find("three"); bucket_index == -1 {
		t.Errorf("Failed to find 3 in the hash table")
	}
}

func TestChainingHashTable_Remove(t *testing.T) {
	t.Parallel()

	cht := hash.NewChainingHashTable[int](10)

	cht.Set("one", 1)
	cht.Set("two", 2)
	cht.Set("three", 3)

	cht.Delete("one")

	if _, bucket_index := cht.Find("one"); bucket_index != -1 {
		t.Errorf("Failed to find 1 in the hash table")
	}

	if _, bucket_index := cht.Find("two"); bucket_index == -1 {
		t.Errorf("Failed to find 2 in the hash table")
	}

	if _, bucket_index := cht.Find("three"); bucket_index == -1 {
		t.Errorf("Failed to find 3 in the hash table")
	}
}

func TestChainingHashTable_Clear(t *testing.T) {
	t.Parallel()

	cht := hash.NewChainingHashTable[int](10)

	cht.Set("one", 1)
	cht.Set("two", 2)
	cht.Set("three", 3)

	cht.Clear()

	if _, bucket_index := cht.Find("one"); bucket_index != -1 {
		t.Errorf("Failed to find 1 in the hash table")
	}

	if _, bucket_index := cht.Find("two"); bucket_index != -1 {
		t.Errorf("Failed to find 2 in the hash table")
	}

	if _, bucket_index := cht.Find("three"); bucket_index != -1 {
		t.Errorf("Failed to find 3 in the hash table")
	}
}

func TestChainingHashTable_Find(t *testing.T) {
	t.Parallel()

	cht := hash.NewChainingHashTable[int](10)

	cht.Set("one", 1)
	cht.Set("two", 2)
	cht.Set("three", 3)

	if _, bucket_index := cht.Find("one"); bucket_index == -1 {
		t.Errorf("Failed to find 1 in the hash table")
	}

	if _, bucket_index := cht.Find("two"); bucket_index == -1 {
		t.Errorf("Failed to find 2 in the hash table")
	}

	if _, bucket_index := cht.Find("three"); bucket_index == -1 {
		t.Errorf("Failed to find 3 in the hash table")
	}
}

func TestChainingHashTable_FindWithStruct(t *testing.T) {
	t.Parallel()

	type Employee struct {
		Name  string
		Phone string
	}

	cht := hash.NewChainingHashTable[Employee](10)

	a := Employee{"John Doe", "555-1234"}
	b := Employee{"Jane Doe", "555-6789"}
	c := Employee{"John Smith", "555-0000"}

	cht.Set(a.Name, a)
	cht.Set(b.Name, b)
	cht.Set(c.Name, c)

	if _, bucket_index := cht.Find(a.Name); bucket_index == -1 {
		t.Errorf("Failed to find 1 in the hash table")
	}

	if _, bucket_index := cht.Find(b.Name); bucket_index == -1 {
		t.Errorf("Failed to find 2 in the hash table")
	}

	if _, bucket_index := cht.Find(c.Name); bucket_index == -1 {
		t.Errorf("Failed to find 3 in the hash table")
	}
}

func TestChainingHashTable_Get(t *testing.T) {
	t.Parallel()

	cht := hash.NewChainingHashTable[int](10)

	cht.Set("one", 1)
	cht.Set("two", 2)
	cht.Set("three", 3)

	if cht.Get("one") != 1 {
		t.Errorf("Failed to get 1 from the hash table")
	}

	if cht.Get("two") != 2 {
		t.Errorf("Failed to get 2 from the hash table")
	}

	if cht.Get("three") != 3 {
		t.Errorf("Failed to get 3 from the hash table")
	}
}

func TestChainingHashTable_Contains(t *testing.T) {
	t.Parallel()

	cht := hash.NewChainingHashTable[int](10)

	cht.Set("one", 1)
	cht.Set("two", 2)
	cht.Set("three", 3)

	if !cht.Contains("one") {
		t.Errorf("Failed to find 1 in the hash table")
	}

	if !cht.Contains("two") {
		t.Errorf("Failed to find 2 in the hash table")
	}

	if !cht.Contains("three") {
		t.Errorf("Failed to find 3 in the hash table")
	}
}
