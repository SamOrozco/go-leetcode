package main

func sort(arr *[]int) {
	mergeSort(arr, make([]int, len(*arr)), 0, len(*arr)-1)
}

func mergeSort(arr *[]int, temp []int, left, right int) {
	if right <= left {
		return
	}
	middle := (right + left) / 2
	mergeSort(arr, temp, left, middle)
	mergeSort(arr, temp, middle+1, right)
	mergeHalves(arr, *arr, temp, left, right)
}

func mergeHalves(arr *[]int, values []int, temp []int, left, right int) {
	leftEnd := (left + right) / 2
	rightEnd := right
	leftIdx := left
	rightIdx := leftEnd + 1
	resultIndex := left
	for leftIdx <= leftEnd && rightIdx <= rightEnd {
		if values[leftIdx] <= values[rightIdx] {
			temp[resultIndex] = values[leftIdx]
			leftIdx++
		} else {
			temp[resultIndex] = values[rightIdx]
			rightIdx++
		}
		resultIndex++
	}

	// still have items in right half
	if rightIdx < rightEnd {
		for i := rightIdx; i <= rightEnd; i++ {
			temp[resultIndex] = values[i]
			resultIndex++
		}
	} else {
		for i := leftIdx; i <= leftEnd; i++ {
			temp[resultIndex] = values[i]
			resultIndex++
		}
	}
	*arr = temp
}

func main() {
	values := make([]int, 3)
	values[0] = 7
	values[1] = 2
	values[2] = 4
	sort(&values)
	for _, val := range values {
		println(val)
	}
}
