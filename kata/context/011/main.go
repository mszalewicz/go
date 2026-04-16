package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	response := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	go func(ctx context.Context, response chan string) {
		randTime := time.Duration(rand.IntN(100)+50) * time.Millisecond
		time.Sleep(randTime)

		select {
		case <-ctx.Done():
			response <- "time out"
		default:
			response <- "3rd party info"
		}
	}(ctx, response)

	select {
	case msg := <-response:
		fmt.Println(msg)
	}
}
