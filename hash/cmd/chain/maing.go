package main

import (
	"fmt"

	"github.com/tankcdr/hash"
)

type Employee struct {
	Name  string
	Phone string
}

func main() {
	employeeHashTable := hash.NewChainingHashTable[string](10)

	// Make some names.
	employees := []Employee{
		Employee{"Ann Archer", "202-555-0101"},
		Employee{"Bob Baker", "202-555-0102"},
		Employee{"Cindy Cant", "202-555-0103"},
		Employee{"Dan Deever", "202-555-0104"},
		Employee{"Edwina Eager", "202-555-0105"},
		Employee{"Fred Franklin", "202-555-0106"},
		Employee{"Gina Gable", "202-555-0107"},
		Employee{"Herb Henshaw", "202-555-0108"},
		Employee{"Ida Iverson", "202-555-0109"},
		Employee{"Jeb Jacobs", "202-555-0110"},
	}

	for _, employee := range employees {
		employeeHashTable.Set(employee.Name, employee.Phone)
	}
	employeeHashTable.Dump()

	fmt.Printf("Table contains Sally Owens: %t\n", employeeHashTable.Contains("Sally Owens"))
	fmt.Printf("Table contains Dan Deever: %t\n", employeeHashTable.Contains("Dan Deever"))
	fmt.Println("Deleting Dan Deever")
	employeeHashTable.Delete("Dan Deever")
	fmt.Printf("Table contains Dan Deever: %t\n", employeeHashTable.Contains("Dan Deever"))
	fmt.Printf("Sally Owens: %s\n", employeeHashTable.Get("Sally Owens"))
	fmt.Printf("Fred Franklin: %s\n", employeeHashTable.Get("Fred Franklin"))
	fmt.Println("Changing Fred Franklin")
	employeeHashTable.Set("Fred Franklin", "202-555-0100")
	fmt.Printf("Fred Franklin: %s\n", employeeHashTable.Get("Fred Franklin"))

	employeeHashTable.Dump()
}
