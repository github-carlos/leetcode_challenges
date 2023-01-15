package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	
	answer := []string{}

	for i := 1; i <= n; i++ {
		if i <=2 {
			answer = append(answer, strconv.Itoa(i))
			continue
		}

		if (i % 3 == 0 && i % 5 == 0) {
			answer = append(answer, "FizzBuzz")
			continue
		}

		if (i % 3 == 0) {
			answer = append(answer, "Fizz")
			continue
		}

		if (i % 5 == 0) {
			answer = append(answer, "Buzz")
			continue
		}

		answer = append(answer, strconv.Itoa(i))
	}

	return answer 
}

func main() {
	fmt.Println(fizzBuzz(3)) // shouuld be ["1", "2", "Fizz"]
}