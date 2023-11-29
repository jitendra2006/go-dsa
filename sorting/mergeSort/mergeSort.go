package main

import "fmt"

func main() {
	inp := []int{10, 3, 4, 1, 7}
	sorted := mergeSort(inp)
	fmt.Println("sorted array is: ", sorted)
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	first := mergeSort(arr[:len(arr)/2])
	second := mergeSort(arr[len(arr)/2:])
	return merge(first, second)
}

func merge(arr1, arr2 []int) []int {
	final := []int{}
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			final = append(final, arr1[i])
			i++
		} else {
			final = append(final, arr2[j])
			j++
		}
	}
	for i < len(arr1) {
		final = append(final, arr1[i])
		i++
	}

	for j < len(arr2) {
		final = append(final, arr2[j])
		j++
	}
	return final

}
