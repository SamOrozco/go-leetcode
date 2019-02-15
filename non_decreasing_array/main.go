package main

import "math"

func main() {
	print(checkPossibility([]int{3, 4, 2, 3}))
	print(checkPossibility([]int{4, 2, 3}))
	print(checkPossibility([]int{4, 2, 1}))
	print(checkPossibility([]int{1}))
	print(checkPossibility([]int{1, 2, 3}))
}

// gonna bubble sort, if I have to do more than one operation it's a no go
func checkPossibility(nums []int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}
	if length == 1 {
		return true
	}
	for i := 0; i < length-1; i++ {
		if nums[i] > nums[i+1] {
			if i == 0 {
				nums[i] = math.MinInt32
			} else {
				nums[i] = nums[i-1]
			}
			if !sorted(nums) {
				return false
			} else {
				return true
			}
		}
	}
	return true
}

func sorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

func swap(nums []int, idx1, idx2 int) {
	temp := nums[idx1]
	nums[idx1] = nums[idx2]
	nums[idx2] = temp
}
