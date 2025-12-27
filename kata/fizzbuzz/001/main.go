package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	usage := "Usage: program <number>"

	if len(os.Args) != 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(usage)
		os.Exit(1)
	}

	mapping := map[int]string{
		3: "fizz",
		5: "buzz",
		7: "bong",
	}

	for key, value := range mapping {
		if number%key == 0 {
			fmt.Print(value, " ")
		}
	}

	fmt.Println()
}
