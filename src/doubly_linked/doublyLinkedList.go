package main

import (
	"fmt"
	"strings"
)

type Cell struct {
	data string
	prev *Cell
	next *Cell
}

type DoublyLinkedList struct {
	topSentinel    *Cell
	bottomSentinel *Cell
}

func makeDoublyLinkedList() DoublyLinkedList {
	list := DoublyLinkedList{
		topSentinel:    &Cell{data: "top_sentinel"},
		bottomSentinel: &Cell{data: "bottom_sentinel"},
	}
	list.topSentinel.next = list.bottomSentinel
	list.bottomSentinel.prev = list.topSentinel

	return list
}

// Add the cell <cell> after the cell <me>
func (me *Cell) addAfter(cell *Cell) {
	// I am making an assumption here (to check if valid)
	//  that only the bottom sentinel has .next == nil
	if me.next == nil {
		panic("Cannot add cell after bottom sentinel")
	}
	cell.next = me.next
	cell.prev = me
	me.next.prev = cell
	me.next = cell

}

// Add the cell <cell> before the cell <me>
func (me *Cell) addBefore(cell *Cell) {
	// I am making an assumption here (to check if valid)
	//  that only the top sentinel has .prev == nil
	if me.prev == nil {
		panic("Cannot add cell before top sentinel")
	}
	cell.prev = me.prev
	cell.next = me
	me.prev.next = cell
	me.prev = cell
}

func (list *DoublyLinkedList) addRange(values []string) {
	// for doubly-linked lists, the last cell
	// is the cell just befopre the bottom sentinel
	var lastCell *Cell = list.bottomSentinel.prev

	// insert the values in
	for _, value := range values {
		var newCell = &Cell{data: value}
		lastCell.next = newCell
		newCell.prev = lastCell
		lastCell = newCell
	}
	// and put the links in to/from the bottomSentinel
	lastCell.next = list.bottomSentinel
	list.bottomSentinel.prev = lastCell

}

func (list *DoublyLinkedList) toString(sep string) string {
	var b strings.Builder
	sep = strings.Trim(sep, " ")
	for curr := list.topSentinel; curr != nil; curr = curr.next {
		if curr.next != nil {
			fmt.Fprintf(&b, "%s %s ", curr.data, sep)
		} else {
			fmt.Fprintf(&b, "%s", curr.data)
		}
	}
	return b.String()
}

func (list *DoublyLinkedList) toStringMax(sep string, max int) string {
	var b strings.Builder
	var ctr int
	sep = strings.Trim(sep, " ")
	for curr := list.topSentinel; curr != nil; curr = curr.next {
		ctr = ctr + 1
		if ctr == max {
			break
		}
		if curr.next != nil {
			fmt.Fprintf(&b, "%s %s ", curr.data, sep)
		} else {
			fmt.Fprintf(&b, "%s", curr.data)
		}
	}
	return b.String()
}
