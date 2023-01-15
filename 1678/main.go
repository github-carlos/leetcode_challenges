package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	
	var maxValue = 0

	for _, accounts := range accounts {
		valueAccounts := 0
		for _, valueAccount := range accounts {
			valueAccounts += valueAccount
		}
		if valueAccounts > maxValue {
			maxValue = valueAccounts
		}
	}

	return maxValue
}

func main() {
	test1 := [][]int{{1, 2, 3}, {3, 2, 1}} // would be 6
	test2 := [][]int{{1, 5}, {7, 3}, {3, 5}} // would be 10
	fmt.Println(maximumWealth(test1))
	fmt.Println(maximumWealth(test2))
	// maximumWealth(test2)
}