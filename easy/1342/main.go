package main

import "fmt"

func numberOfSteps(num int) int {

	if num == 0 {
		return 0
	}

	nOfSteps := 1

	for num > 1 {
		nOfSteps++
		if num & 1 == 0 {
			num = num / 2
			continue
		}
		num -=1
	}

	return nOfSteps
}

func main() {
	fmt.Println(numberOfSteps(14)) // should be 6
	fmt.Println(numberOfSteps(8)) // should be 4
	fmt.Println(numberOfSteps(123)) // should be 12
}