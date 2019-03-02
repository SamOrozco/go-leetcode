package main

func main() {
	print(checkPossibility([]int{3, 4, 2, 3}))
	print(checkPossibility([]int{4, 2, 3}))
	print(checkPossibility([]int{4, 2, 1}))
	print(checkPossibility([]int{1}))
	print(checkPossibility([]int{1, 2, 3}))
	print(checkPossibility([]int{2, 3, 3, 2, 4}))
}

// gonna bubble sort, if I have to do more than one operation it's a no go
func checkPossibility(nums []int) bool {
	numLength := len(nums)
	if numLength < 1 {
		return true
	}
	lastVal := nums[0]
	for i := 1; i < numLength; i++ {
		currentVal := nums[i]
		if currentVal < lastVal {
			// set last value in array to current value
			// check if sorted
			nums[i-1] = currentVal
			if !isSorted(nums) {
				if i >= 2 {
					// rest change
					nums[i-1] = lastVal
					// copy previous value
					nums[i] = nums[i-1]
					return isSorted(nums)
				}
			} else {
				return true
			}
		}
		lastVal = currentVal
	}
	return true
}

func isSorted(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			return true
		}
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
