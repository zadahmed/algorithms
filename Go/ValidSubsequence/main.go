// Algorithm to validate if subsequence within a sequence

package main

// Time o(n) / Space o(n)
func validSubsequence(array []int, subsequence []int) bool {
	arrayIndex := 0
	subseqIndex := 0
	for arrayIndex < len(array) && subseqIndex < len(subsequence) {
		if array[arrayIndex] == subsequence[subseqIndex] {
			subseqIndex += 1
		}
		arrayIndex += 1
	}
	return subseqIndex == len(subsequence)
}
