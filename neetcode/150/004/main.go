package main

import "fmt"

// Given an array of strings strs, group all anagrams together into sublists. You may return the output in any order.

// An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

// Example 1:

// Input: strs = ["act","pots","tops","cat","stop","hat"]

// Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]

// Example 2:

// Input: strs = ["x"]

// Output: [["x"]]

// Example 3:

// Input: strs = [""]

// Output: [[""]]

// Constraints:

//     1 <= strs.length <= 1000.
//     0 <= strs[i].length <= 100
//     strs[i] is made up of lowercase English letters.

func groupAnagrams(strs []string) [][]string {
    mapping := make(map[[26]int][]string)

    for _, str := range strs {
        var frequency [26]int

        for _, char := range str {
            frequency[int(char) - int('a')] += 1
        }

        mapping[frequency] = append(mapping[frequency], str)
    }

    answer := make([][]string, 0, len(mapping))

    for _, group := range mapping {
        answer = append(answer, group)
    }

    return answer
}

func main() {
    strs := []string{"act","pots","tops","cat","stop","hat"}

    fmt.Println(groupAnagrams(strs))

}