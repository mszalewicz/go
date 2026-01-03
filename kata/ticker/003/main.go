package main

import (
	"context"
	"time"
)

func cleanup(ticker *time.Ticker) {
	println("Stopping ticker")
	ticker.Stop()
}

func main() {
	ticker := time.NewTicker(150 * time.Millisecond)
	defer cleanup(ticker)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

forever:
	for {
		select {
		case <-ctx.Done():
			break forever
		case <-ticker.C:
			println(time.Now().String())
		}
	}
}
