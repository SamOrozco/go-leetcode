package main

import "os"

func main() {
	if trailingZeroes(7) != 1 {
		println("Failed")
		os.Exit(1)
	}

	if trailingZeroes(30) != 7 {
		println("Failed")
		os.Exit(1)
	}
}

func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}
	fact := factorial(int64(n))
	if fact < 10 {
		return 0
	}

	numZero := 0
	for fact > 0 {
		curVal := popTrailInt(fact)
		if curVal == 0 {
			numZero++
		} else {
			return numZero
		}
		fact /= 10
	}

	return numZero

}

func popTrailInt(val int64) int64 {
	return val % 10
}

func factorial(n int64) int64 {
	value := n
	n -= 1
	for n > 0 {
		value *= n
		n--
	}
	return value
}
