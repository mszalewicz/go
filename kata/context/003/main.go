package main

import (
	"context"
	"fmt"
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
		fmt.Println("Query timed out.")
	case message := <-response:
		fmt.Println(message)
	}
}

func mockThirdPartyQuery(ctx context.Context, response chan string) {
	rndTime := time.Duration(rand.UintN(300)) * time.Millisecond
	time.Sleep(rndTime + 150*time.Millisecond)

	select {
	case <-ctx.Done():
		return
	default:
		response <- "3rd party data {...}"
	}
}
