package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abba"
	fmt.Println(isPalindrome(str))
}

func isPalindrome(s string) bool {
	rmSpecialChar := strip([]byte(s))
	str := string(rmSpecialChar)
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ToLower(str)

	if len(str)%2 != 0 {
		for l, r := len(str)/2, len(str)/2; r < len(str); {
			if str[l] != str[r] {
				return false
			}
			l--
			r++
		}
	}

	for l, r := 0, len(str)-1; l < r; {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func strip(s []byte) []byte {
	n := 0
	for _, b := range s {
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			s[n] = b
			n++
		}
	}
	return s[:n]
}
