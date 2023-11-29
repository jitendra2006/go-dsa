package main

import "fmt"

func main() {
	inp := []int{10, 1, 12, 2, 33, 45, 5, -3, 0}
	selectionSort(inp)
	fmt.Println("sorted array: ", inp)
}

func selectionSort(arr []int) {
	// [10,1,12,2,33,45,5,-3,0]
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		temp := arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
}
