package main

import (
	"fmt"
	"strings"
)

type Cell[T any] struct {
	data T
	prev *Cell[T]
	next *Cell[T]
}

type DoublyLinkedList[T any] struct {
	topSentinel    *Cell[T]
	bottomSentinel *Cell[T]
}

func makeDoublyLinkedList[T any](topValue T, bottomValue T) DoublyLinkedList[T] {
	list := DoublyLinkedList[T]{
		topSentinel:    &Cell[T]{data: topValue},
		bottomSentinel: &Cell[T]{data: bottomValue},
	}
	list.topSentinel.next = list.bottomSentinel
	list.bottomSentinel.prev = list.topSentinel

	return list
}

// Add the cell <cell> after the cell <me>
func (me *Cell[T]) addAfter(cell *Cell[T]) {
	if me.next == nil {
		panic("Cannot add cell after bottom sentinel")
	}
	cell.next = me.next
	cell.prev = me
	me.next.prev = cell
	me.next = cell

}

// Add the cell <cell> before the cell <me>
func (me *Cell[T]) addBefore(cell *Cell[T]) {
	if me.prev == nil {
		panic("Cannot add cell before top sentinel")
	}
	cell.prev = me.prev
	cell.next = me
	me.prev.next = cell
	me.prev = cell
}

func (list *DoublyLinkedList[T]) addRange(values []T) {
	// for doubly-linked lists, the last cell
	// is the cell just before the bottom sentinel
	var lastCell *Cell[T] = list.bottomSentinel.prev

	// insert the values in
	for _, value := range values {
		var newCell = &Cell[T]{data: value}
		lastCell.next = newCell
		newCell.prev = lastCell
		lastCell = newCell
	}
	// and put the links in to/from the bottomSentinel
	lastCell.next = list.bottomSentinel
	list.bottomSentinel.prev = lastCell

}

func (list *DoublyLinkedList[T]) toString(sep string) string {
	var b strings.Builder
	sep = strings.Trim(sep, " ")
	for curr := list.topSentinel; curr != nil; curr = curr.next {
		if curr.next != nil {
			fmt.Fprintf(&b, "%v %v ", curr.data, sep)
		} else {
			fmt.Fprintf(&b, "%v", curr.data)
		}
	}
	return b.String()
}

func (list *DoublyLinkedList[T]) toStringMax(sep string, max int) string {
	var b strings.Builder
	var ctr int
	sep = strings.Trim(sep, " ")
	for curr := list.topSentinel; curr != nil; curr = curr.next {
		ctr = ctr + 1
		if ctr == max {
			break
		}
		if curr.next != nil {
			fmt.Fprintf(&b, "%v %v ", curr.data, sep)
		} else {
			fmt.Fprintf(&b, "%v", curr.data)
		}
	}
	return b.String()
}

type Queue[T any] struct {
	items DoublyLinkedList[T]
}

func makeQueue[T any](topValue T, bottomValue T) Queue[T] {
	var queue Queue[T]
	queue.items = makeDoublyLinkedList[T](topValue, bottomValue)
	return queue
}

// since the queue.item.topSentinel is accessed quite often
// define a function to access it
func (queue *Queue[T]) top() *Cell[T] {
	return queue.items.topSentinel
}

// since the queue.item.bottomSentinel is accessed quite often
// define a function to access it
func (queue *Queue[T]) bottom() *Cell[T] {
	return queue.items.bottomSentinel
}

func (queue *Queue[T]) enqueue(item T) {
	newCell := &Cell[T]{data: item}
	queue.bottom().addBefore(newCell)
}

func (queue *Queue[T]) dequeue() T {

	// if the queue is empty, calling dequeue() will panic
	if queue.isEmpty() {
		panic("Dequeueing an empty queue")
	}
	// top := queue.items.topSentinel
	var removedCell = queue.top().next
	var newFirstCell = removedCell.next

	queue.top().next = removedCell.next
	newFirstCell.prev = queue.top()

	return removedCell.data
}

func (queue *Queue[T]) isEmpty() bool {
	// queue is empty if all there is are the top and bottom sentinels
	// and they are pointing to each other
	top := queue.top()
	bottom := queue.bottom()

	return top.next == bottom && bottom.prev == top
}
func (queue *Queue[T]) toString(sep string) string {
	return queue.items.toString(sep)
}

// test driver
func main() {
	var queue = makeQueue[string]("top", "bottom")
	queue.enqueue("1")
	queue.enqueue("2")
	queue.enqueue("3")
	fmt.Println(queue.toString(">"))
}