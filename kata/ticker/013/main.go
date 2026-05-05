package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	ticker := time.NewTicker(30 * time.Millisecond)

	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			fmt.Println("done")
			return
		case <-ticker.C:
			fmt.Println(time.Now())
		}
	}
}
