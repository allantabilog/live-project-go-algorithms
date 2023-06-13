package main

import (
	"fmt"
	"strings"
)

type Cell struct {
	data string
	next *Cell
}

type LinkedList struct {
	sentinel *Cell
}

// returns an empty linked list
func makeLinkedList() LinkedList{
	// return LinkedList{ sentinel: nil } // my original code
	return LinkedList {
		sentinel: &Cell{data: "Sentinel", next: nil},
	}
}

// Add the cell <cell> after the cell <me>
func (me *Cell) addAfter(cell *Cell) {
	cell.next = me.next
	me.next = cell
}
func (me *Cell) deleteAfter() *Cell {
	deletedCell := me.next 
	if deletedCell == nil {
		panic("No cell to delete")
	}
	// point me to whatever the deleted cell was pointing to
	me.next = deletedCell.next
	return deletedCell
}

// add a list of values to the list
// in the same order as the given list
func (list *LinkedList) addRange(values []string) {
	var lastCell *Cell 

	// follow the 'next' pointers to locate the end of the list
	for curr := list.sentinel; curr != nil; curr = curr.next {
		lastCell = curr
	}

	// lastCelll should now point to the end of the list
	// insert the values in
	for _, value := range values {
		var newCell = &Cell{data: value}
		lastCell.next = newCell
		lastCell = newCell 
	}

}

// representation of the linked list as a string
// takes a separator character e.g. "->"
// list looks like:
// Sentinel -> 1 -> 2 -> 3
func (list *LinkedList) toString(sep string) string {
	var b strings.Builder
	sep = strings.Trim(sep, " ");
	for curr := list.sentinel; curr != nil; curr = curr.next {
		if curr.next != nil {
			fmt.Fprintf(&b, "%s %s ", curr.data, sep)
		} else {
			fmt.Fprintf(&b, "%s", curr.data)
		}
	}
	return b.String()
}

// Return an array containing the list's contents
// (excludes the sentinel)
func (list *LinkedList) toArray() []string {
	var contents []string
	for curr := list.sentinel.next; curr != nil; curr = curr.next {
		contents = append(contents, curr.data)
	}
	return contents
}

// Returns the length of the list == number of contents
func (list *LinkedList) length() int {
	var length int = -1

	for curr := list.sentinel; curr != nil; curr = curr.next {
		length += 1
	}

	return length
}

// True iff the list has no contents
func (list *LinkedList) isEmpty() bool {
	return list.sentinel.next == nil
}