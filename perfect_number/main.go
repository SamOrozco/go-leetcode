package main

func main() {
	print(checkPerfectNumber(28))
}

func checkPerfectNumber(num int) bool {
	if num < 1 {
		return false
	}
	return getDivisorSum(num) == num
}

func getDivisorSum(num int) int {
	root := num / 2
	divisors := root
	for i := 1; i < int(root); i++ {
		if num%i == 0 {
			divisors += i
		}
	}
	return divisors
}
