package main

import (
	"os"
)

func main() {
	val0 := []int{3, 1, 4, 1, 5}
	if findPairs(val0, 2) != 2 {
		println("Fail")
		os.Exit(1)
	}

	val1 := []int{1, 2, 3, 4, 5}
	if findPairs(val1, 1) != 4 {
		println("Fail")
		os.Exit(1)
	}

	val2 := []int{1, 3, 1, 5, 4}
	if findPairs(val2, 0) != 1 {
		println("Fail")
		os.Exit(1)
	}
}

func findPairs(nums []int, k int) int {
	//public int findPairs(int[] nums, int k) {
	//	if (k < 0) { return 0; }
	//
	//	Set<Integer> starters = new HashSet<Integer>();
	//	Set<Integer> uniqs = new HashSet<Integer>();
	//	for (int i = 0; i < nums.length; i++) {
	//		if (uniqs.contains(nums[i] - k)) starters.add(nums[i] - k);
	//		if (uniqs.contains(nums[i] + k)) starters.add(nums[i]);
	//		uniqs.add(nums[i]);
	//	}
	//
	//	return starters.size();
	//}
	if k < 0 {
		return 0
	}

	start := make(map[int]bool)
	unq := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := unq[nums[i]-k]; ok {
			start[nums[i]-k] = true
		}

		if _, ok := unq[nums[i]+k]; ok {
			start[nums[i]] = true
		}
		unq[nums[i]] = true
	}
	return len(start)
}
