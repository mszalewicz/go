package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Case struct {
	denominator int
	text string
}

func main() {
    usage := "Usage:\nprogram <number>\n"

    if len(os.Args) != 2 {
        log.Fatal(usage)
    }

    number, err := strconv.Atoi(os.Args[1])

    if err != nil {
        log.Fatal(usage)
    }

    mapping := []Case {
       {3, "fizz"},
       {5, "buzz"},
       {7, "back"},
    }

    for i := range mapping {
        if number % mapping[i].denominator == 0 {
            fmt.Println(mapping[i].text)
        }
    }
}
