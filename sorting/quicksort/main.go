package main

import "fmt"

func main() {
	nums := []int{9, 7, 4, 3, 1, 2, 5, 6, 8}
	fmt.Println(nums)
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func findKthLargest(nums []int, k int) int {
	idxToFind := len(nums) - k
	quickSort(nums, 0, len(nums)-1)
	return nums[idxToFind]
}

func quickSort(nums []int, left int, right int) {
	if left < right {
		// get partition
		partitionIdx := partition(nums, left, right)
		// quick sort left side
		quickSort(nums, left, partitionIdx-1)
		// quick sort right side
		quickSort(nums, partitionIdx+1, right)
	}
}

func partition(nums []int, left int, right int) int {
	pivotElem := nums[right]
	partitionIdx := left

	for j := left; j < right; j++ {
		if nums[j] < pivotElem {
			// Swap position i and j
			nums[partitionIdx], nums[j] = nums[j], nums[partitionIdx]
			partitionIdx++
		}
	}
	nums[partitionIdx], nums[right] = nums[right], nums[partitionIdx]
	return partitionIdx
}
