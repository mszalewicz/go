package main

import (
	"strconv"
	"testing"
)

func TestCheckNumber(t *testing.T) {
	name := "TestCheckNumber "

	cases := []struct {
		input int
		expected string
	}{
		{3, "fizz "},
		{5, "buzz "},
		{7, "bring "},
		{35, "buzz bring "},
		{105, "fizz buzz bring a"},
	}

	for index, test := range cases {
		t.Run(name + strconv.Itoa(index), func(t *testing.T) {
			result := checkNumber(test.input)
			if test.expected != result {
				t.Errorf("Expected: '%s', got: '%s'", test.expected, result)
			}
		})
	}
}
