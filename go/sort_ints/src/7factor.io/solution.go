package main

import "fmt"

func main() {

	fmt.Println("the go kata")

	arr := makeRandomArray(Range{0, 10000})

	fmt.Println(arr, "will sort %v")

	sorted, err := sortInts(arr)

	if err != nil {
		_ = fmt.Errorf("caught error while running sort: %v", err)
	}

	fmt.Println(sorted, "sorted array is: %v")
}

func sortInts(arr []int) ([]int, error) {
	if arr == nil {
		return make([]int, 0), nil
	}

	if len(arr) <= 1 {
		return arr, nil
	}

	builder := make([]int, 0)
	first := arr[0]
	lessThan := make([]int, 0)
	greaterThan := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		if arr[i] < first {
			lessThan = append(lessThan, arr[i])
		}
		if arr[i] > first {
			greaterThan = append(greaterThan, arr[i])
		}
	}

	lessThan, _ = sortInts(lessThan)
	builder = append(builder, lessThan...)
	builder = append(builder, first)
	greaterThan, _ = sortInts(greaterThan)
	builder = append(builder, greaterThan...)

	return builder, nil
}
