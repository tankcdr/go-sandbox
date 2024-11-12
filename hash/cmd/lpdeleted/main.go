package main

import (
	"fmt"
	"time"

	"math/rand"

	"github.com/tankcdr/hash"
)

type Employee struct {
	Name  string
	Phone string
}

func main() {
	// Make some names.
	employees := []Employee{
		{"Ann Archer", "202-555-0101"},
		{"Bob Baker", "202-555-0102"},
		{"Cindy Cant", "202-555-0103"},
		{"Dan Deever", "202-555-0104"},
		{"Edwina Eager", "202-555-0105"},
		{"Fred Franklin", "202-555-0106"},
		{"Gina Gable", "202-555-0107"},
	}

	hashTable := hash.NewLinearProbingHashTable[Employee](10)
	for _, employee := range employees {
		hashTable.Set(employee.Name, employee)
		println("Setting", employee.Name)
	}
	hashTable.Dump()

	hashTable.Probe("Hank Hardy")
	fmt.Printf("Table Contains Sally Owens: %t\n", hashTable.Contains("Sally Owens"))
	fmt.Printf("Table Contains Dan Deever: %t\n", hashTable.Contains("Dan Deever"))
	fmt.Println("Deleting Dan Deever")
	hashTable.Delete("Dan Deever")
	fmt.Printf("Table Contains Dan Deever: %t\n", hashTable.Contains("Dan Deever"))
	fmt.Printf("Sally Owens: %s\n", hashTable.Get("Sally Owens"))
	fmt.Printf("Fred Franklin: %s\n", hashTable.Get("Fred Franklin"))
	fmt.Println("Changing Fred Franklin")
	hashTable.Set("Fred Franklin", Employee{"Fred Franklin", "202-555-0100"})
	fmt.Printf("Fred Franklin: %s\n", hashTable.Get("Fred Franklin"))
	hashTable.Dump()

	hashTable.Probe("Ann Archer")
	hashTable.Probe("Bob Baker")
	hashTable.Probe("Cindy Cant")
	hashTable.Probe("Dan Deever")
	hashTable.Probe("Edwina Eager")
	hashTable.Probe("Fred Franklin")
	hashTable.Probe("Gina Gable")
	hashTable.Set("Hank Hardy", Employee{"Fred Franklin", "202-555-0100"})
	hashTable.Probe("Hank Hardy")

	// Look at clustering.
	fmt.Println(time.Now())                   // Print the time so it will compile if we use a fixed seed.
	random := rand.New(rand.NewSource(12345)) // Initialize with a fixed seed
	// random := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed
	bigCapacity := 1009
	bigHashTable := hash.NewLinearProbingHashTable[string](bigCapacity)
	numItems := int(float32(bigCapacity) * 0.9)

	key := "899-156610"

	for i := 0; i < numItems; i++ {
		str := fmt.Sprintf("%d-%d", i, random.Intn(1000000))
		if i == 500 {
			key = str
		}
		bigHashTable.Set(str, str)
	}

	bigHashTable.Delete(key)

	bigHashTable.DumpConcise()
	fmt.Printf("Average Probe sequence length: %f\n",
		bigHashTable.AveProbeSequenceLength())
}
