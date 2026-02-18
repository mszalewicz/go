package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	usage := "Usage: program <number>"

	if len(os.Args) != 2 {
		log.Fatal(usage)
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(usage)
	}

	result := checkNumber(number)

	fmt.Println(result)
}

func checkNumber(number int) string {
	var sb strings.Builder

	mapping := map[int]string{
		3: "fizz",
		5: "buzz",
		7: "bring",
	}

	for key, value := range mapping {
		if number%key == 0 {
			sb.WriteString(value)
			sb.WriteString(" ")
		}
	}

	return sb.String()
}
