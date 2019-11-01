package main

import "fmt"

func main() {
	fmt.Println("the go kata")
}

func bubbleSort(arr []int) ([]int, error) {
	if arr == nil {
		return make([]int, 0), nil
	}

	if len(arr) <= 1 {
		return arr, nil
	}

	if arr[0] > arr[1] {
		tmp := arr[1]
		arr[1] = arr[0]
		arr[0] = tmp
	}

	fmt.Println(arr)

	return arr, nil
}
