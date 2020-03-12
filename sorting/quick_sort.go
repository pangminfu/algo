package sorting

func QuickSort(arr []int) []int {
	low := 0
	high := len(arr) - 1

	quickSort(arr, low, high)

	return arr
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]

	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arri := arr[j]
			arr[j] = arr[i]
			arr[i] = arri
		}
	}

	arr[high] = arr[i+1]
	arr[i+1] = pivot

	return i + 1
}
