package main

import "fmt"

func main() {
	inp := []int{22, 12, 10, 3, 6, 19, 15, -1, -5, -11}
	bubbleSort(inp)
	fmt.Println("sorted array: ", inp)
}

// for optimizing the bubble sort we can use a flag to indicate that swapping is occured to check that array is not sorted initially.
func bubbleSort(arr []int) {

	for step := 0; step < len(arr)-1; step++ {
		swapped := 0
		for i := 0; i < len(arr)-step-1; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
				swapped++
			}
		}
		if swapped == 0 {
			break
		}
	}
}
