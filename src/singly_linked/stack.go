package main

import "fmt"

// Represent a Stack as a struct
// with one member: <contents> which is a linked list
type Stack struct {
	contents LinkedList
}

// Returns an empty stack
func makeStack() Stack {
	return Stack {
		contents: makeLinkedList(),
	}
}

// True if the stack is empty
func (s *Stack) isEmpty() bool {
	return s.contents.isEmpty()
}

// push an item to the top of the stack
func (s *Stack) push(item string) {
	s.contents.sentinel.addAfter(&Cell{data: item})
}

// pop the item on the top of the stack
// panics if an empty stack is popped
// returns the popped item otherwise
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