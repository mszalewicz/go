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

	sorted := make([][]string, 0, len(mapping))

	for _, entry := range mapping {
		sorted = append(sorted, entry)
	}

	return sorted
}

func main() {
	strs := []string{"test", "estt", "new", "ewn", "bas"}

	fmt.Println(categories(strs))
}
