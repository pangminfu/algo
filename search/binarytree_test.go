package search_test

import (
	"search"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	testCases := []struct {
		name           string
		arr            []int
		target         int
		expectedResult int
	}{
		{
			name: "",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			search.BinaryTree(testCase.arr, testCase.target)
		})
	}
}
