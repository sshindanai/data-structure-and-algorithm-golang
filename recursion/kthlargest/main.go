package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}

	fmt.Println(findKthLargest(arr, 4))

}

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
