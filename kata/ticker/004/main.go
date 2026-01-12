package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now())
		case <-ctx.Done():
			fmt.Println("finished")
			return
		}
	}
}
