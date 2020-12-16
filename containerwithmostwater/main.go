package main

import (
	"fmt"
)

func main() {
	height := []int{1, 2, 4, 3}
	fmt.Println(maxAreaBrtueForce(height))
	fmt.Println(maxArea(height))
}

// Optimal
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	max := 0

	for i, j := 0, len(height)-1; i < j; {
		area := (j - i) * min(height[i], height[j])

		if area > max {
			max = area
		}

		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}

	return max
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// brute force
func maxAreaBrtueForce(height []int) int {
	var maxVal int
	for i := 0; i < len(height); i++ {
		for j := len(height) - 1; j >= i+1; j-- {
			width := j - i
			var heightArea int

			if height[i] <= height[j] {
				heightArea = height[i]
			} else {
				heightArea = height[j]
			}
			area := width * heightArea
			if area > maxVal {
				maxVal = area
			}
		}
	}
	return maxVal
}

// func maxArea(height []int) int {

// }

// Test case
// Input: height = [1,1]
// Output: 1

// Example 3:

// Input: height = [4,3,2,1,4]
// Output: 16

// Example 4:

// Input: height = [1,2,1]
// Output: 2
