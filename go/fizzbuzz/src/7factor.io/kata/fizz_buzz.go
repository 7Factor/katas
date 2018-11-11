package kata

import (
	"strconv"
)

func FizzBuzz(min int, max int) []string {
	var output []string

	for i := min; i <= max; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			output = append(output, "FizzBuzz")
		} else if i % 3 == 0 {
			output = append(output, "Fizz")
		} else if i % 5 == 0 {
			output = append(output, "Buzz")
		} else {
			output = append(output, strconv.Itoa(i))
		}
	}
	return output
}
