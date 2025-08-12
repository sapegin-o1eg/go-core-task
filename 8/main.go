package main

import (
	"fmt"
	"task8/utils"
	"time"
)

func main() {
	fmt.Println("Custom WaitGroup Example")

	wg := utils.NewCustomWaitGroup()

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, wg)
	}

	fmt.Println("Waiting for all workers to complete...")
	wg.Wait()
	fmt.Println("All workers completed!")
}

func worker(id int, wg *utils.CustomWaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting...\n", id)
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Worker %d completed!\n", id)
}
