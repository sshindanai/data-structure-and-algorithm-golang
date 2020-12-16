package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abccdba"
	fmt.Println(validPalindrome(str))
}

func check(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func validPalindrome(s string) bool {
	s = strings.ToLower(s)

	for i, j := 0, len(s)-1; i <= j; {
		if check(s[i]) == false {
			i++
			continue
		}
		if check(s[j]) == false {
			j--
			continue
		}
		if s[i] != s[j] {
			return validSubPalindrome(s, i+1, j) || validSubPalindrome(s, i, j-1)
		}
		i++
		j--
	}
	return true
}

func validSubPalindrome(s string, l int, r int) bool {
	for l < r {
		if string(s[l]) != string(s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}
