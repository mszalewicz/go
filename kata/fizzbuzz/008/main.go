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
		return
	}

	number, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(usage)
		return
	}	

	mapping := map[int]string{
		3: "fizz",
		5: "buzz",
		7: "ding",
	}

	for denominator, text := range mapping {
		if number % denominator == 0 {
			fmt.Println(text)
		}
	}
}
