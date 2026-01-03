package main

import (
	"context"
	"math/rand/v2"
	"time"
)

func main() {
	response := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	go mockThirdPartyQuery(ctx, response)

	select {
	case <-ctx.Done():
		println("Query timed out")
	case msg := <-response:
		println(msg)
	}
}

func mockThirdPartyQuery(ctx context.Context, response chan string) {
	rndTime := time.Duration(rand.UintN(300)) * time.Millisecond
	padTime := 150 * time.Millisecond

	time.Sleep(rndTime + padTime)

	select {
	case <-ctx.Done():
		return
	default:
		response <- "3rd party data {...}"
	}
}
