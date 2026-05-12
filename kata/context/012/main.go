package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	response := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	go thirdparty(ctx, response)

	select {
	case <-ctx.Done():
		fmt.Println("timed out")
	case msg := <-response:
		fmt.Println(msg)
	}
}

func thirdparty(ctx context.Context, response chan string) {
	rndTime := time.Duration(rand.IntN(220)+100) * time.Millisecond
	fmt.Println(rndTime)
	time.Sleep(rndTime)

	select {
	case <-ctx.Done():
		return
	default:
		response <- "third party info..."
	}
}
