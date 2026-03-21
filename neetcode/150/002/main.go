package main

import (
	"fmt"
)

// Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.
//
// An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.
//
// Example 1:
//
// Input: s = "racecar", t = "carrace"
//
// Output: true
//
// Example 2:
//
// Input: s = "jar", t = "jam"
//
// Output: false
//
// Constraints:
//
//     s and t consist of lowercase English letters.

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	a := make([]int, 26)
	b := make([]int, 26)

	for i := range len(s) {
		a[s[i]-byte('a')] += 1
		b[t[i]-byte('a')] += 1
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func main() {
	s := "racecar"
	t := "carrace"

	fmt.Println(isAnagram(s, t))
}
