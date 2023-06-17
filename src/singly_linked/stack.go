package main

import "fmt"

// Represent a Stack as a struct
// with one member: <contents> which is a linked list
type Stack struct {
	items LinkedList
}

// Returns an empty stack
func makeStack() Stack {
	return Stack{
		items: makeLinkedList(),
	}
}

// True if the stack is empty
func (s *Stack) isEmpty() bool {
	return s.items.isEmpty()
}

// push an item to the top of the stack
func (s *Stack) push(item string) {
	s.items.sentinel.addAfter(&Cell{data: item})
}

// pop the item on the top of the stack
// panics if an empty stack is popped
// returns the popped item otherwise
func (s *Stack) pop() string {
	if s.isEmpty() {
		panic("Cannot pop an empty stack")
	}

	poppedItem := s.items.sentinel.next.data
	s.items.sentinel.deleteAfter()
	return poppedItem
}

func (s *Stack) length() int {
	return s.items.length()
}

func (s *Stack) toString() string {
	return s.items.toString("->")
}

func (s *Stack) toArray() []string {
	return s.items.toArray()
}

func (s *Stack) peek() {
	fmt.Printf("Stack contents:%v\n", s.items.toString(">"))
}
