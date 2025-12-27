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
		fmt.Println("Third party query timed out.")
		return
	case responseBody := <-response:
		fmt.Println(responseBody)
	}
}

func mockThirdPartyQuery(ctx context.Context, response chan string) {
	rndTimeout := time.Duration(rand.Int32N(300) + 150) * time.Millisecond
	time.Sleep(rndTimeout)

	select {
	case <-ctx.Done():
		return
	default:
		response <- "Third party data {...}"
	}
}
