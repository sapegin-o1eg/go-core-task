package main

import (
	"fmt"
	"task5/utils"
)

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	ok, result := utils.Intersection(a, b)
	fmt.Printf("Intersection of %+v and %+v: %+v (%t)\n", a, b, result, ok)

	ok, result = utils.Intersection(a, []int{})
	fmt.Printf("Intersection of %+v and %+v: %+v (%t)\n", a, []int{}, result, ok)

	ok, result = utils.Intersection([]int{}, b)
	fmt.Printf("Intersection of %+v and %+v: %+v (%t)\n", []int{}, b, result, ok)

	ok, result = utils.Intersection([]int{}, []int{})
	fmt.Printf("Intersection of %+v and %+v: %+v (%t)\n", []int{}, []int{}, result, ok)
}
