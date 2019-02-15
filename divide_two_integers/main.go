package main

import (
	"math"
	"os"
)

func main() {
	if val := divide(10, 3); val != 3 {
		print("fail")
		os.Exit(1)
	}
	if val := divide(7, -3); val != -2 {
		print("fail")
		os.Exit(1)
	}

	if val := divide(-1<<31, -1); val != 1<<31-1 {
		print("fail")
		os.Exit(1)
	}
	println("pass")
}

func divide(dividend int, divisor int) int {
	if divisor == 1 {
		return dividend
	}
	if dividend == divisor {
		return 1
	}
	if abs(dividend) == abs(divisor) {
		return -1
	}
	// takes care of overflow
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	neg := divisor < 0
	if neg {
		divisor *= -1
	}

	cnt := 0
	temp := dividend
	for temp > divisor {
		temp -= divisor
		cnt++
	}

	if neg {
		return cnt * -1
	}
	return cnt
}

func abs(val int) int {
	if val < 0 {
		return - val
	}
	return val
}
