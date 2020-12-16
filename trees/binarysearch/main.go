package main

import (
	"math"
)

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8

}

func binarySearch(arr []int, l int, r int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	for l <= r {
		mid := math.Floor(float64(l)+float64(r)) / 2
		found := arr[int(mid)]

		if found == target {
			return int(mid)
		} else if target < found {
			r = int(mid) - 1
		} else if target > found {
			l = int(mid) + 1
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	firstPos := binarySearch(nums, 0, len(nums)-1, target)
	if firstPos == -1 {
		return []int{-1, -1}
	}
}
