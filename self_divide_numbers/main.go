package main

func main() {
	res := selfDividingNumbers(1, 22)
	for _, e := range res {
		print(e)
	}

}

func selfDividingNumbers(left int, right int) []int {
	result := make([]int, 0)
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			result = append(result, i)
		}
	}
	return result
}

func isSelfDividing(num int) bool {
	orig := num
	for num > 0 {
		digit := num % 10
		if digit == 0 {
			return false
		}
		num /= 10
		if orig%digit != 0 {
			return false
		}
	}
	return true
}
