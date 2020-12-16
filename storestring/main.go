package main

import "fmt"

func main() {
	fmt.Println(storeString("aaxbbf", "ab"))
}

func storeString(s string, target string) int {
	dict := make(map[string]int)

	for i := 0; i < len(s); i++ {
		val := string(s[i])
		dict[val]++
	}
	counter := 0

	for i := 0; true; i++ {
		char := string(target[i])

		if v, ok := dict[char]; ok && v != 0 {

		}
	}

	return counter
}
