package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 3, 7, 9, 2}
	target := 11
	//fmt.Println(TwoSumBruteForce(arr, target))
	fmt.Println(twoSum(arr, target))
}

func twoSum(arr []int, target int) []int {
	if len(arr) == 0 || len(arr) == 1 {
		return nil
	}

	m := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		numberToFind := target - arr[i]
		if v, ok := m[numberToFind]; ok {
			return []int{v, i}
		}
		m[arr[i]] = i
	}
	return nil
}

func TwoSumBruteForce(arr []int, target int) []int {
	if len(arr) == 0 || len(arr) == 1 {
		return []int{}
	}
	for l := 0; l < len(arr); l++ {
		numToFind := target - arr[l]

		for r := l + 1; r < len(arr); r++ {
			if numToFind == arr[r] {
				return []int{arr[l], arr[r]}
			}
		}
	}

	return []int{}
}
