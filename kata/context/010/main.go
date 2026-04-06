package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	msgBroker := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	go func(ctx context.Context, msgBroker chan string) {
		randTime := time.Duration(rand.IntN(100))*time.Millisecond + 150*time.Millisecond
		time.Sleep(randTime)

		select {
		case <-ctx.Done():
			return
		default:
			msgBroker <- "3rd party info: ..."
		}
	}(ctx, msgBroker)

	select {
	case <-ctx.Done():
		fmt.Println("- query timed out -")
	case msg := <-msgBroker:
		fmt.Println(msg)
	}
}
