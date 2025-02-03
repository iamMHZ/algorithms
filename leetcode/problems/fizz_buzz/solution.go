package fizz_buzz

import (
	"strconv"
)

func FizzBuzz(n int) []string {

	var array []string = make([]string, n)

	for i := 0; i < n; i++ {
		number := i + 1
		if number%3 == 0 && number%5 == 0 {
			array[i] = "FizzBuzz"
		} else if number%3 == 0 {
			array[i] = "Fizz"
		} else if number%5 == 0 {
			array[i] = "Buzz"
		} else {
			array[i] = strconv.Itoa(number)
		}
	}

	return array

}
