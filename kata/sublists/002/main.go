package main

import "fmt"

func categories(strs []string) [][]string {
    mapping := make(map[[26]int][]string)

    for _, str := range strs {
        var frequency [26]int
        for _, char := range str {
            frequency[int(char)-int('a')] += 1
        }

        mapping[frequency] = append(mapping[frequency], str)
    }

    categories := make([][]string, 0, len(mapping))

    for _, strs := range mapping {
        categories = append(categories, strs)
    }

    return categories
}

func main() {
    strs := []string{"test", "estt", "bee"}

    fmt.Println(categories(strs))
}
