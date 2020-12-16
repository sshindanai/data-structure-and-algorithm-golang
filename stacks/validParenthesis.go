package main

// var parenthesis = map[string]string{
// 	"(": ")",
// 	"{": "}",
// 	"[": "]",
// }

// func isValid(s string) bool {
// 	if len(s) == 0 {
// 		return true
// 	}
// 	if len(s)%2 != 0 {
// 		return false
// 	}

// 	var stack Stack

// 	for i := 0; i < len(s); i++ {
// 		if _, ok := parenthesis[string(s[i])]; ok {
// 			stack.Push(string(s[i]))
// 		} else {
// 			leftBracket := stack.Peek()
// 			correctBracket := parenthesis[leftBracket]

// 			if string(s[i]) != correctBracket {
// 				return false
// 			}
// 		}
// 	}

// 	return stack.IsEmpty()

// }
