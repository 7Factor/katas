package kata

import (
	"fmt"
	"reflect"
	"testing"
)

func findPrime(n int) []int {
	if n < 2 {
		return []int{}
	}
	if n == 2 {
		return []int{2}
	}
	if n % 2 == 0 {
		if n % 3 == 0 {
			return []int{2, 3}
		}
		if n % 4 == 0 && n >= 8 {
			return []int{2,2,2}
		}
		return []int{n/2, n/2}
	}
	return []int{n}
}

type tests []struct {
	given int
	want  []int
}

func TestKata(t *testing.T) {
	tests := tests{
		{1, []int{}},
		{2, []int{2}},
		{3, []int{3}},
		{4, []int{2, 2}},
		{5, []int{5}},
		{6, []int{2,3}},
		{7, []int{7}},
		{8, []int{2, 2, 2}},
	}
	runTests(tests, t)
}

func runTests(tests tests, t *testing.T) {
	for _, test := range tests {
		testname := fmt.Sprintf("%d", test.given)
		t.Run(testname, func(t *testing.T) {
			ans := findPrime(test.given)
			if !reflect.DeepEqual(ans, test.want) {
				t.Errorf("result was %d, wanted %d", ans, test.want)
			}
		})
	}
}