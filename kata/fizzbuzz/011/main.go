package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const USAGE = "Usage: program <number>"

func main() {
	if len(os.Args) != 2 {
		log.Fatal(USAGE)
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Error during number conversion. Message -> ", err)
	}

	fmt.Println(checkNumber(number))
}

func checkNumber(number int) string {
	mapping := map[int]string{
		3: "fizz",
		5: "buzz",
		7: "ayo",
	}

	var builder strings.Builder
	keyNumber := 1

	for key, value := range mapping {
		if number%key == 0 {
			if keyNumber > 1 {
				builder.WriteString(" ")
			}

			builder.WriteString(value)

			keyNumber++
		}

	}

	return builder.String()
}
