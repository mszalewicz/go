package main

import "fmt"

func main() {
	strs := []string{"test", "estt", "map", "amp", "crow"}

	fmt.Println(sublists(strs))
}

func sublists(strs []string) [][]string {
	lists := make(map[[26]int][]string)

	for _, str := range strs {
		var shape [26]int

		for _, char := range str {
			shape[int(char)-'a']++
		}

		lists[shape] = append(lists[shape], str)
	}

	result := make([][]string, 0, len(lists))
	for _, group := range lists {
		result = append(result, group)
	}

	return result
}
