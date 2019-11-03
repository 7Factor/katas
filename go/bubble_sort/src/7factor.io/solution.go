package main

import "fmt"

func main() {
	fmt.Println("the go kata")
}

type ArrayAndIndex struct {
	arr   []int
	index int
}

func bubbleSort(arr []int) ([]int, error) {
	if arr == nil {
		return make([]int, 0), nil
	}

	if len(arr) <= 1 {
		return arr, nil
	}

	for i := 0; i < len(arr)-1; i++ {
		compareAndSwap(ArrayAndIndex{arr: arr, index:i})
		for j := 0; j < len(arr)-1; j++ {
			compareAndSwap(ArrayAndIndex{arr: arr, index:j})
		}
	}

	fmt.Println(arr)

	return arr, nil
}

func compareAndSwap(ai ArrayAndIndex) {
	if ai.arr[ai.index] > ai.arr[ai.index+1] {
		tmp := ai.arr[ai.index+1]
		ai.arr[ai.index+1] = ai.arr[ai.index]
		ai.arr[ai.index] = tmp
	}
}
