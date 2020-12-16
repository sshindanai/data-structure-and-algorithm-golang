package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(tailFact(26))
}

func factorial(n int) int {
	if n < 1 {
		return 1
	}

	return n * factorial(n-1)
}

// Tail factorial implement =====
func tailFact(n int) int {
	return factT(n-1, n)
}

func factT(n, current int) int {
	if n == 1 {
		return current
	}
	return factT(n-1, n*current)
}

// ===============================

func countDown(num int) {
	time.Sleep(time.Second * 1)
	if num == 0 {
		fmt.Println("BOOOM!!!")
	} else {
		fmt.Println(num)
		countDown(num - 1)
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func reverseString(s []byte) {
	l := 0
	r := len(s) - 1

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
