package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(10 * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			fmt.Println("stopped")
			return
		case <-ticker.C:
			fmt.Println(time.Now())
		}
	}
}
