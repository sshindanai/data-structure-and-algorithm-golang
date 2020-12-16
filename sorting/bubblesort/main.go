package main

import "fmt"

func main() {
	arr := []int{5, 3, 6, 7, 8, 4, 2}
	fmt.Println(bubbleSort(arr))
}

func bubbleSort(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}

func selectionSort(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}
