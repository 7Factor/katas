package kata

import (
	"strconv"
)

func FizzBuzz(min int, max int) []string {
	var output []string

	for i := min; i <= max; i++ {
		output = append(output, Output(i))
	}

	return output
}

func Output(value int) string {
	var output string
	if value % 3 == 0 && value % 5 == 0 {
		output = "FizzBuzz"
	} else if value % 3 == 0 {
		output = "Fizz"
	} else if value % 5 == 0 {
		output = "Buzz"
	} else {
		output = strconv.Itoa(value)
	}
	return output
}
