package search_test

import (
	"errors"
	"testing"

	"github.com/pangminfu/algo/search"
	"github.com/stretchr/testify/assert"
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
		{
			name: "Given valid arr of length 5 taget is located at index 0",
			arr: []int{
				2,
				5,
				7,
				8,
				15,
			},
			target:         2,
			expectedResult: 0,
			expectedErr:    nil,
		},
		{
			name: "Given valid arr of length 5 taget is located at last index",
			arr: []int{
				2,
				5,
				7,
				8,
				15,
			},
			target:         15,
			expectedResult: 4,
			expectedErr:    nil,
		},
		{
			name: "Given valid arr of length 5 but target is not found",
			arr: []int{
				2,
				5,
				7,
				8,
				15,
			},
			target:         16,
			expectedResult: -1,
			expectedErr:    errors.New("not found"),
		},
		{
			name: "Given valid arr of length 1",
			arr: []int{
				2,
			},
			target:         2,
			expectedResult: 0,
			expectedErr:    nil,
		},
		{
			name: "Given valid arr of length 1 but target not found",
			arr: []int{
				2,
			},
			target:         3,
			expectedResult: -1,
			expectedErr:    errors.New("not found"),
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
