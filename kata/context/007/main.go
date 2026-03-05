package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"	
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Millisecond)
	defer cancel()

	response := make(chan string)

	// Third party query
	go func(ctx context.Context, response chan string) {
		rndTime :=	time.Duration(rand.Int32N(70) + 50) * time.Millisecond
		time.Sleep(rndTime)	

		select {
		case <-ctx.Done():
			return
		default:
			response <-"3rd party data"
		}
	}(ctx, response)
	
	select {
	case <-ctx.Done():
		fmt.Println("query timed out")
		return
	case msg := <-response:
		fmt.Println(msg)
	}
}
