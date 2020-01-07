package main

import (
	"fmt"
)

type queue []int

func (s *queue) push(number int) { *s = append(*s, number) }
func (s *queue) pop() int {
	firstNumber := []int(*s)[1]
	if len(*s) == 1 {
		*s = queue{}
	} else {
		*s = []int(*s)[1:]
	}
	return firstNumber
}

func main() {
	s := new(queue)

	s.push(10)
	s.push(20)
	s.push(30)
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
}
