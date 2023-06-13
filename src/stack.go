package main

import "fmt"

type Stack struct {
	contents LinkedList
}

func makeStack() Stack {
	return Stack {
		contents: makeLinkedList(),
	}
}

func (s *Stack) isEmpty() bool {
	return s.contents.isEmpty()
}

func (s *Stack) push(item string) {
	s.contents.sentinel.addAfter(&Cell{data: item})
}

func (s *Stack) pop() string {
	if s.isEmpty() {
		panic("Cannot pop an empty stack")
	}
	
	poppedItem := s.contents.sentinel.next.data
	s.contents.sentinel.deleteAfter()
	return poppedItem
}

func (s *Stack) length() int {
	return s.contents.length()
}

func (s *Stack) toString() string{
	return s.contents.toString("->")
}

func (s *Stack) toArray() []string {
	return s.contents.toArray()
}

func (s *Stack) peek() {
	fmt.Printf("Stack contents:%v\n", s.contents.toString(">"))
}