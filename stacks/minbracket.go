package main

// func minRemoveToMakeValid(s string) string {
// 	if len(s) == 0 {
// 		return ""
// 	}

// 	var stack Stack

// 	slice := strings.Split(s, "")

// 	for i := 0; i < len(slice); i++ {
// 		if slice[i] == "(" {
// 			stack.Push(slice[i])
// 		} else if slice[i] == ")" && !stack.IsEmpty() && stack.Peek() == "(" {
// 			stack.Pop()
// 		} else if slice[i] == ")" {
// 			slice[i] = ""
// 		}
// 	}

// 	return strings.Join(slice, "")
// }
