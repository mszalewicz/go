package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	response := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()

	go thirdPartyQuery(ctx, response)

	select {
	case <-ctx.Done():
		fmt.Println("query timed out")
	case msg := <-response:
		fmt.Println(msg)
	}
}

func thirdPartyQuery(ctx context.Context, response chan string) {
	rndTime := 100*time.Millisecond + time.Duration(rand.IntN(100))*time.Millisecond
	time.Sleep(rndTime)

	select {
	case <-ctx.Done():
		return
	default:
		response <- "third party data: {...}"
	}
}
