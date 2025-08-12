package main

import (
	"fmt"
	"task7/utils"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	merged := utils.MergeChannels(ch1, ch2, ch3)

	ch1 <- 1
	ch2 <- 2
	close(ch1)
	close(ch2)

	go func() {
		defer close(ch3)
		delay := 500 * time.Millisecond
		for i := 3; i <= 10; i++ {
			time.Sleep(delay)
			ch3 <- i
			delay /= 2
		}
	}()

	for num := range merged {
		fmt.Println(num)
	}
}
