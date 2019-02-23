package main

import "os"

func main() {
	if shortestSubarray([]int{17, 85, 93, -45, -21}, 150) != 2 {
		println("failed")
		os.Exit(1)
	}

	if shortestSubarray([]int{1}, 1) != 1 {
		println("failed")
		os.Exit(1)
	}

	if shortestSubarray([]int{1, 2}, 4) != -1 {
		println("failed")
		os.Exit(1)
	}

	if val := shortestSubarray([]int{2, -1, 2}, 3); val != 3 {
		println("failed")
		os.Exit(1)
	}

	if val := shortestSubarray([]int{27, 20, 79, 87, -36, 78, 76, 72, 50, -26}, 453); val != 3 {
		println("failed")
		os.Exit(1)
	}
}

const badValue = -1

func shortestSubarray(A []int, K int) int {
	aLength := len(A)
	if aLength%2 == 0 {
		aLength -= 1
	}
	return sumUntilFound(A, K, make([]int, aLength), aLength, 1)
}

func sumUntilFound(haystack []int, needle int, sumSpace []int, sumSize int, count int) int {
	if len(haystack) == 1 {
		if haystack[0] == needle {
			return count
		} else {
			return -1
		}
	}
	haystackLength := len(haystack)
	if haystackLength < 1 {
		return badValue
	}

	idx := 0
	for idx = 0; idx < sumSize; idx++ {
		currentVal := haystack[idx]
		if currentVal == needle {
			return count
		}
		if idx != (haystackLength - 1) {
			nextVal := haystack[idx+1]
			sum := currentVal + nextVal
			if sum >= needle { // we found the needle
				return count + 1
			}
			sumSpace[idx] = sum
		} else {
			lastVal := currentVal
			if lastVal == needle {
				return count
			}
			sumSpace[idx] = lastVal
		}
	}

	if sumSize%2 != 0 {
		sumSpace[idx] = haystack[idx]
	}

	return sumUntilFound(sumSpace, needle, make([]int, sumSize-1), sumSize-1, count+1)
}
