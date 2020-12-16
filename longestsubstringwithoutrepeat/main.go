package main

import "fmt"

func main() {
	str := "abccabb"
	fmt.Println(lengthOfLongestSubstring(str))

}

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	seen := make(map[string]int)
	longestSubStr := 0

	for l, r := 0, 0; r < len(s); r++ {
		currChar := string(s[r])
		prevSeenChar := seen[currChar]

		if _, ok := seen[currChar]; ok {
			if prevSeenChar >= l {
				l = prevSeenChar + 1
			}
		}
		seen[currChar] = r

		longestSubStr = max(longestSubStr, r-l+1)
	}

	return longestSubStr
}

// Brute force solution
func lengthOfLongestSubstringBF(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	longestSubStr := 0

	for l := 0; l < len(s); l++ {
		seen := make(map[string]bool)
		currLength := 0

		for r := l; r < len(s); r++ {
			if !seen[string(s[r])] {
				currLength++
				seen[string(s[r])] = true
				longestSubStr = max(longestSubStr, currLength)
			} else {
				break
			}
		}

	}

	return longestSubStr
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
