package main

import (
	"fmt"
)

type stack []int

func (s *stack) push(number int) { *s = append(*s, number) }
func (s *stack) pop() int {
	lastNumber := []int(*s)[len(*s)-1]
	*s = []int(*s)[:len(*s)-1]
	return lastNumber
}

func main() {
	s := new(stack)

	s.push(10)
	s.push(20)
	s.push(30)
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
}
