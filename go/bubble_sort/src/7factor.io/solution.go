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

	if len(arr) == 2 {
		if arr[0] > arr[1] {
			tmp := arr[1]
			arr[1] = arr[0]
			arr[0] = tmp
		}
		return arr, nil
	}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			tmp := arr[i+1]
			arr[i+1] = arr[i]
			arr[i] = tmp
		}
	}

	fmt.Println(arr)

	return arr, nil
}
