package main

import (
	"fmt"
)

func main() {
    strs := []string{"test", "sett", "ball", "lalb", "nope"}

    fmt.Println(categories(strs))
}

func categories(strs []string) [][]string {
    categories := make(map[[26]int][]string)

    for _, str := range strs {
        var frequency [26]int

        for _, char := range str {
            frequency[int(char) - int('a')] += 1
        }

        categories[frequency] = append(categories[frequency], str)
    }

    result := make([][]string, 0, len(categories))

    for _, category := range categories {
        result = append(result, category)
    }

    return result
}