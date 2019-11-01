package main

import "fmt"

func main() {
	fmt.Println("the go kata")
}

func bubbleSort(arr []int) ([]int, error) {
	if arr == nil {
		return make([]int, 0), nil
	}
	return arr, nil
}
