package main

import (
	"fmt"
	"task3/utils"
)

func main() {
	sm1 := utils.NewStringIntMap()

	// Different output formats:
	fmt.Printf("[sm1] Standard output: %v\n", sm1)
	fmt.Printf("[sm1] Without pointer (dereference): %v\n", *sm1)
	fmt.Printf("[sm1] Only map content: %v\n", sm1.M)
	fmt.Printf("[sm1] Pretty output: %+v\n", sm1)
	fmt.Printf("[sm1] Detailed output: %#v\n", sm1)

	sm1.Add("a", 1)
	fmt.Printf("[sm1] After adding element: %+v\n", sm1.M)
	sm1.Add("b", 2)
	sm1.Add("c", 3)
	sm1.Add("d", 4)
	sm1.Add("e", 5)
	sm1.Add("f", 6)
	sm1.Add("g", 7)
	sm1.Add("h", 8)
	sm1.Add("i", 9)
	sm1.Add("j", 10)
	sm1.Add("k", 11)
	fmt.Printf("[sm1] After adding elements: %+v\n", sm1.M)

	sm1.Remove("k")
	fmt.Printf("[sm1] After deleting key k: %+v\n", sm1.M)

	sm2 := sm1.Copy()
	fmt.Printf("[sm2] Copy of sm1: %+v\n", sm2)

	fmt.Printf("[sm1] Does key 'k' exist in sm1?: %v\n", sm1.Exists("k"))
	fmt.Printf("[sm1] Does key 'a' exist in sm1?: %v\n", sm1.Exists("a"))

	value1, ok1 := sm1.Get("a")
	fmt.Printf("[sm1] Value of key 'a': %v, and it exists: %v\n", value1, ok1)

	value2, ok2 := sm1.Get("k")
	fmt.Printf("[sm1] Value of key 'k': %v, and it exists: %v\n", value2, ok2)
}
