package main

import "fmt"

func categorise(strs []string) [][]string {
    mapping := make(map[[26]int][]string)

    for _, str := range strs {
        var frequency [26]int

        for _, char := range str {
            frequency[int(char) - int('a')] += 1
        }

        mapping[frequency] = append(mapping[frequency], str)
    }

    categories := make([][]string, 0, len(mapping))

    for _, entry := range mapping {
        categories = append(categories, entry)
    }

    return categories
}

func main() {
    strs := []string{"act","pots","tops","cat","stop","hat", "hole", "leho"}

    fmt.Println(categorise(strs))
}