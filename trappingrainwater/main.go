package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trapBF(height))
	fmt.Println(trap(height))
}

func trap(height []int) int {
	maxLeft, maxRight := 0, 0
	totalWater := 0
	for l, r := 0, len(height)-1; l <= r; {
		if height[l] <= height[r] {
			if height[l] >= maxLeft {
				maxLeft = height[l]
			} else {
				totalWater += maxLeft - height[l]
			}
			l++
		} else {
			if height[r] >= maxRight {
				maxRight = height[r]
			} else {
				totalWater += maxRight - height[r]
			}
			r--
		}
	}
	return totalWater
}

// bruteforce
func trapBF(height []int) int {
	total := 0
	for i := 0; i < len(height); i++ {
		leftP, rightP := i, i
		maxLeft, maxRight := 0, 0

		for leftP >= 0 {
			maxLeft = max(maxLeft, height[leftP])
			leftP--
		}

		for rightP < len(height) {
			maxRight = max(maxRight, height[rightP])
			rightP++
		}
		currenWater := min(maxLeft, maxRight) - height[i]

		if currenWater >= 0 {
			total += currenWater
		}
	}
	return total
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
