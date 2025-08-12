package main

import (
	"context"
	"fmt"
	"task6/utils"
)

func main() {
	randCh := utils.RandomGenerator(context.Background(), 10, 1, 10)
	i := 0
	fmt.Printf("Let's generate 10 random numbers from 1 to 10:\n")
	for num := range randCh {
		fmt.Printf("Num #%d: %d\n", i, num)
		i++
	}
}
