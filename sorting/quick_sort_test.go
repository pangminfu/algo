package sorting_test

import (
	"testing"

	"github.com/pangminfu/algo/sorting"
	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		name           string
		arr            []int
		expectedResult []int
		expectedErr    error
	}{
		{
			name: "Given arr",
			arr: []int{
				10,
				80,
				30,
				90,
				40,
				50,
				70,
			},
			expectedResult: []int{
				10,
				30,
				40,
				50,
				70,
				80,
				90,
			},
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := sorting.QuickSort(testCase.arr)

			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}
