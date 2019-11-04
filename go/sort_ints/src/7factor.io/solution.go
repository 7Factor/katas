package main

import "fmt"

func main() {
	fmt.Println("the go kata")
}

func sortInts(arr []int) ([]int, error) {
	if arr == nil {
		return make([]int, 0), nil
	}

	if len(arr) <= 1 {
		return arr, nil
	}

	first := arr[0]
	second := arr[1]
	if first < second {
		arr = []int{first, second}
	} else {
		arr = []int{second, first}
	}

	fmt.Println(arr)

	return arr, nil
}
