package main

func main() {
	println(smallestRangeII([]int{0, 10}, 2))
	println(smallestRangeII([]int{1}, 0))
	println(smallestRangeII([]int{1, 3, 6}, 3))
}

func smallestRangeII(A []int, K int) int {
	if K == 0 {
		return 0
	}
	length := len(A)
	highVals := make([]int, length)
	lowVals := make([]int, length)
	for i := 0; i < length; i++ {
		high := A[i] + K
		low := A[i] - K
		highVals[i] = high
		lowVals[i] = low
	}

	//target := 0
	currentSum := 0
	for i := 0; i < length; i++ {
		temphigh := currentSum + highVals[i]
		tempLow := currentSum + lowVals[i]
		if temphigh < tempLow {
			currentSum = temphigh
		} else {
			currentSum = tempLow
		}
	}
	return currentSum
}

func diff(left, right int) int {
	if left > right {
		return left - right
	}
	return right - left
}

//Given an array A of integers, for each integer A[i] we need to choose either x = -K or x = K, and add x to A[i] (only once).
//
//After this process, we have some array B.
//
//Return the smallest possible difference between the maximum value of B and the minimum value of B.
