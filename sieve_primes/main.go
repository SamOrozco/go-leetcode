package main

func main() {
	vals := sievesPrimes(1000)
	for _, i := range vals {
		println(i)
	}
	print(search(vals, 653))
}

func search(vals []int, val int) int {
	lo := 0
	hi := len(vals)
	for lo <= hi {
		var mid = lo + (hi-lo)/2
		var midValue = vals[mid]
		if midValue == val {
			return mid
		} else if midValue > val {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1

}

func sievesPrimes(size int) []int {
	seen := make([]bool, size)
	primes := make([]int, 0)
	for i := 2; i < size; i++ {
		if val := seen[i]; val {
			continue
		} else {
			primes = append(primes, i)
			seen[i] = true
			for j := i; j < size; j += i {
				seen[j] = true
			}
		}
	}
	return primes
}
