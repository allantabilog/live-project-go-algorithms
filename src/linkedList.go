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
func main() {

	listOne := makeLinkedList()
	fmt.Println(listOne);

}	
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
	me.next = nil 
	return deletedCell
}

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

func (list *LinkedList) toString(sep string) string {
	var b strings.Builder
	for curr := list.sentinel; curr != nil; curr = curr.next {
		if curr.next != nil {
			fmt.Fprintf(&b, "%s %s", curr.data, sep)
		} else {
			fmt.Fprintf(&b, "%s", curr.data)
		}
	}
	return b.String()
}