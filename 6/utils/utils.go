package utils

import (
	"context"
	"math/rand"
	"time"
)

// RandomGenerator generates random numbers and sends them to a channel
func RandomGenerator(ctx context.Context, count, min, max int) <-chan int {
	ch := make(chan int)
	
	go func() {
		defer close(ch)
		
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		
		for i := 0; i < count; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- r.Intn(max-min+1) + min:
			}
		}
	}()
	
	return ch
}