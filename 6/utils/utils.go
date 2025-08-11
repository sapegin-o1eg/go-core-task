package utils

import (
	"context"
	"math/rand"
	"time"
)

func RandomGenerator(ctx context.Context, count int, min, max int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		rand.New(rand.NewSource(time.Now().UnixNano()))

		for i := 0; i < count; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				num := rand.Intn(max-min+1) + min

				select {
				case out <- num:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	return out
}
