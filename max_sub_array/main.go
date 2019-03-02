package main

import "sort"

func main() {
	println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(A []int) (int) {
	aLength := len(A)
	sums := make([]int, 0)
	for i := 0; i < aLength; i++ {
		curVal := A[i]
		//
		sums = addSums(curVal, sums)
		// currentVal is a sum
		sums = append(sums, curVal)
	}

	sort.Ints(sums)
	return sums[0]
}

func addSums(val int, sums []int) []int {
}
