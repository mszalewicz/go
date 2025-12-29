package main

import (
	"strconv"
	"testing"
)

func TestCheckNumber(t *testing.T) {
	name := "CheckNumber"

	cases := []struct {
		input    int
		expected string
	}{
		{input: 3, expected: "fizz"},
		{input: 5, expected: "buzz"},
		{input: 7, expected: "ayo"},
		{input: 105, expected: "fizz buzz ayo"},
	}

	for i, test := range cases {
		t.Run(name+strconv.Itoa(i+1), func(t *testing.T) {
			result := checkNumber(test.input)

			if test.expected != result {
				t.Errorf("Expected: '%s', got: '%s'", test.expected, result)
			}
		})
	}
}
