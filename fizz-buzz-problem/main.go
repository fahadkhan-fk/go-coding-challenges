// Problem Statement
// You need to print numbers from 1 to n, but with the following conditions:

// If a number is divisible by 3, print "Fizz" instead of the number.
// If a number is divisible by 5, print "Buzz" instead of the number.
// If a number is divisible by both 3 and 5 (i.e., 15), print "FizzBuzz" instead of the number.
// If the number is not divisible by 3 or 5, print the number itself.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	output := fizzBuzz(15)
	fmt.Println("result: ", output)
}

// Time Complexity O(n) ; Space Complexity O(n)
func fizzBuzz(n int) []string {
	var output []string
	for i := 1; i <= n; i++ {
		str := ""

		if i%3 == 0 {
			str += "Fizz"
		}

		if i%5 == 0 {
			str += "Buzz"
		}

		if str == "" {
			// strconv.Itoa used for integer to string conversion
			str += strconv.Itoa(i)
		}
		output = append(output, str)
	}
	return output
}

// Alternative clean approach
// func fizzBuzz(n int) []string {
// 	arr := []string{}
// 	for i := 1; i <= n; i++ {
// 		if i%3 == 0 && i%5 == 0 {
// 			arr = append(arr, "FizzBuzz")
// 		} else if i%3 == 0 {
// 			arr = append(arr, "Fizz")
// 		} else if i%5 == 0 {
// 			arr = append(arr, "Buzz")
// 		} else {
// 			arr = append(arr, strconv.Itoa(i))
// 		}
// 	}
// 	return arr
// }
