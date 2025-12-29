package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticker := time.NewTicker(50 * time.Millisecond)

	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			fmt.Println(": Done :")
			os.Exit(0)
		case <-ticker.C:
			fmt.Println(time.Now())
		}
	}
}