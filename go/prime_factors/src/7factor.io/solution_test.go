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
	}
	runTests(tests, t)
}

func runTests(tests tests, t *testing.T) {
	for _, test := range tests {
		testname := fmt.Sprintf("%d", test.given)
		t.Run(testname, func(t *testing.T) {
			ans := findPrime(test.given)
			if !reflect.DeepEqual(ans, test.want) {
				t.Errorf("result was %d, wanted []", test.want)
			}
		})
	}
}