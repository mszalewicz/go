package main

import "fmt"

func main() {
	words := []string{"test", "estt", "mob", "bom", "inne"}
	fmt.Println(sublists(words))
}

func sublists(words []string) [][]string {
	groupings := map[[26]int][]string{}

	for _, text := range words {
		var frequency [26]int

		for _, char := range text {
			frequency[int(char)-int('a')]++
		}

		groupings[frequency] = append(groupings[frequency], text)
	}

	result := make([][]string, 0, len(groupings))

	for _, group := range groupings {
		groupedWords := []string{}
		for _, word := range group {
			groupedWords = append(groupedWords, word)
		}
		result = append(result, groupedWords)
	}

	return result
}
