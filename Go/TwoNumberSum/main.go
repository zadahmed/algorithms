// Algorithm that finds pairs of numbers in an array that would sum to a target number provided
package main

import (
	"fmt"
)

// Non efficient solution
// O(n^2) Time / O(1) space complexity
func TwoNumberSum(array []int, targetSum int) []int {
	for i := 0; i < len(array)-1; i++ {
		firstNum := array[i]
		for j := i + 1; j < len(array); j++ {
			secondNum := array[j]
			if firstNum+secondNum == targetSum {
				return []int{firstNum, secondNum}
			}
		}
	}
	return []int{}
}

//efficient solution
//O(n) Time / o(1) space complexity

func TwoNumberSumHashTable(array []int, targetSum int) []int {
	hash_table := map[int]bool{}
	for _, num := range array {
		potentialMatch := targetSum - num
		if _, found := hash_table[potentialMatch]; found {
			return []int{potentialMatch, num}
		} else {
			hash_table[num] = true
		}
	}
	return []int{}
}

func main() {
	Result := TwoNumberSumHashTable([]int{3, 5, -2, 5, 0, 9, 10, 8}, 18)
	fmt.Println(Result)
}
