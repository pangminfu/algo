package search_test

import (
	"testing"

	"github.com/pangminfu/algo/search"
	"gopkg.in/go-playground/assert.v1"
)

func TestBinaryTree(t *testing.T) {
	testCases := []struct {
		name           string
		arr            []int
		target         int
		expectedResult int
		expectedErr    error
	}{
		{
			name: "Given valid arr of length 5",
			arr: []int{
				2,
				5,
				7,
				8,
				15,
			},
			target:         8,
			expectedResult: 3,
			expectedErr:    nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := search.BinaryTree(testCase.arr, testCase.target)

			assert.Equal(t, testCase.expectedErr, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}
