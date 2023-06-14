package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLoop(t *testing.T) {
	
	cellA := Cell{data: "A"}
	cellB := Cell{data: "B"}

	list := makeLinkedList()

	list.sentinel.addAfter(&cellA)
	list.sentinel.next.addAfter(&cellB)

	fmt.Println(list.toString(">"))
	
	// create a loop
	list.sentinel.next.next.next = &cellA
	assert.Equal(t, 
		"Sentinel > A > B > A > B > A > B > A > B > ", 
		list.toStringMax(">", 10),
	)
	
}

func TestHasLoop_Negative(t *testing.T) {

	list := makeLinkedList()
	list.addRange([]string{"1", "2", "3", "4", "5"})

	assert.False(t, list.hasLoop())
}
func TestHasLoop_Positive(t *testing.T) {

	list := makeLinkedList()
	first  := &Cell{data: "first"}
	second := &Cell{data: "second"}
	third := &Cell{data: "third"}

	// create a loop
	list.sentinel.addAfter(first)
	list.sentinel.next.addAfter(second)
	list.sentinel.next.next.addAfter(third)
	list.sentinel.next.next.next.next = first

	assert.Equal(t, 
		"Sentinel > first > second > third > first > second > third > first > second > ",
		list.toStringMax(">", 10),
	)

	assert.True(t, list.hasLoop())

	list2 := makeLinkedList()
	cellA := &Cell{data: "A"}
	cellB := &Cell{data: "B"}

	list2.sentinel.addAfter(cellA)
	list2.sentinel.next.addAfter(cellB)
	list2.sentinel.next.next.next = cellA

	assert.Equal(t, 
		"Sentinel > A > B > A > B > A > B > A > B > ",
		list2.toStringMax(">", 10),
	)

	assert.True(t, list2.hasLoop())




}