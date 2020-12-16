package main

import "fmt"

func main() {
	s := "ab#c"
	t := "ax#C"

	fmt.Println(backspaceCompare(s, t))
}

// brute force
func backspaceCompare(s string, t string) bool {
	const backSpace = "#"
	resultS := []byte{}
	resultT := []byte{}
	counterS := 0
	counterT := 0

	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == backSpace {
			counterS++
			continue
		}

		if counterS > 0 {
			counterS--
			continue
		} else {
			resultS = append(resultS, s[i])
		}
	}

	for i := len(t) - 1; i >= 0; i-- {
		if string(t[i]) == backSpace {
			counterT++
			continue
		}

		if counterT > 0 {
			counterT--
			continue
		} else {
			resultT = append(resultT, t[i])
		}
	}

	return string(resultS) == string(resultT)
}
