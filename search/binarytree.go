package search

import "errors"

func BinaryTree(arr []int, target int) (int, error) {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			return mid, nil
		}
	}

	return -1, errors.New("not found")
}
