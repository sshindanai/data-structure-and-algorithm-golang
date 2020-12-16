package main

import (
	"fmt"
	"strings"
)

func main() {
	str := ")))(((a"

	fmt.Println(minRemoveToMakeValid(str))
}

// Linked list
type ListNode struct {
	Val  string
	Next *ListNode
	Prev *ListNode
}

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
}

type LinkedListQueue LinkedList
type LinkedListStack LinkedList

func (s *LinkedListStack) Push(n *ListNode) {
	if s.Head == nil {
		s.Head = n
		s.Tail = n
		return
	}
	s.Tail.Next = n
	n.Prev = s.Tail
	s.Tail = n
}

func (s *LinkedListStack) Pop() (string, bool) {
	if s.Head == nil {
		return "", false
	}
	elem := s.Tail.Val
	newTail := s.Tail.Prev
	newTail.Next = nil
	s.Tail = newTail

	return elem, true
}

func (s *LinkedListStack) IsEmpty() bool {
	if s.Head == nil {
		return true
	}
	return false
}

func (s LinkedListStack) Print() {
	toPrint := s.Head
	for toPrint != nil {
		fmt.Printf(" %v ", toPrint.Val)
		toPrint = toPrint.Next
	}
	fmt.Printf("\n")
}

func (s LinkedListStack) Peek() string {
	return s.Tail.Val
}

// Slice
type Stack []string
type StackMinBrac []int
type Queue []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// IsEmpty: check if stack is empty
func (s *StackMinBrac) IsEmpty() bool {
	return len(*s) == 0
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Push a new value onto the stack
func (s *StackMinBrac) Push(elem int) {
	*s = append(*s, elem) // Simply append the new value to the end of the stack
}

func (q *Queue) Push(str string) {
	*q = append(*q, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.
	return element, true
}
func (s *StackMinBrac) Pop() int {
	if s.IsEmpty() {
		return 0
	}
	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.
	return element
}
func (q *Queue) Pop() (string, bool) {
	if q.IsEmpty() {
		return "", false
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, true
}

func (s *Stack) Peek() string {
	if s.IsEmpty() {
		return ""
	}
	return (*s)[len(*s)-1]
}
func (s *StackMinBrac) Peek() int {
	if s.IsEmpty() {
		return 0
	}
	return (*s)[len(*s)-1]
}
func (q *Queue) Peek() string {
	if q.IsEmpty() {
		return ""
	}
	return (*q)[0]
}

var parenthesis = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}

	var stack Stack

	for i := 0; i < len(s); i++ {
		if _, ok := parenthesis[string(s[i])]; ok {
			stack.Push(string(s[i]))
		} else {
			leftBracket, _ := stack.Pop()
			correctBracket := parenthesis[leftBracket]

			if string(s[i]) != correctBracket {
				return false
			}
		}
	}

	return stack.IsEmpty()

}

func minRemoveToMakeValid(s string) string {
	slice := []byte(s)
	var stack = []int{}

	for i := 0; i < len(slice); i++ {
		if slice[i] == '(' {
			stack = append(stack, i)
		} else if slice[i] == ')' {
			if len(stack) > 0 {
				// Pop
				lastIndex := len(stack) - 1
				stack = stack[:lastIndex]
			} else {
				slice[i] = '*'
			}
		}
	}

	for _, v := range stack {
		slice[v] = '*'
	}
	res := string(slice)
	str := strings.ReplaceAll(res, "*", "")
	return str
}
