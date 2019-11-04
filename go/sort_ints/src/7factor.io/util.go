package main

type Range struct {
	min int
	max int
}

func makeRandomArray(r Range) []int {
	a := make([]int, r.max-r.min+1)
	for i := range a {
		a[i] = r.min + i
	}
	return a
}
