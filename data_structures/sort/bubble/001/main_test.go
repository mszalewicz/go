package main

import "testing"

func TestBubbleSort(t *testing.T) {
	test := struct {
		input    []int
		expected []int
	}{
		input:    []int{-8422, 9103, -157, 4321, -9980, 765, 5520, -3142, 88, 1204, -6731, 9995, 0, -42, 7311},
		expected: []int{-9980, -8422, -6731, -3142, -157, -42, 0, 88, 765, 1204, 4321, 5520, 7311, 9103, 9995},
	}

	t.Run("default", func(t *testing.T) {
		result := bubbleSort(test.input)

		for i := 0; i < len(test.expected); i++ {
			if test.expected[i] != result[i] {
				t.Errorf("Index: %d, expected value: '%d', got: '%d'", i, test.expected[i], result[i])
				t.Errorf("\nExpected: %v\nGot:      %v", test.expected, result)
				return
			}
		}
	})
}
