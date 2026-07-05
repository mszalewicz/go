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

	go func(response chan string) {
		rndTime := (time.Duration(50) + time.Duration(rand.IntN(100))) * time.Millisecond
		time.Sleep(rndTime)

		response <- "third party data"
	}(response)

	select {
	case <-ctx.Done():
		fmt.Println("canceled")
		return
	case msg := <-response:
		fmt.Println(msg)
	}
}
