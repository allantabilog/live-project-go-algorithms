package main

type Queue struct {
	items DoublyLinkedList
}

func makeQueue() Queue {
	var queue Queue
	queue.items = makeDoublyLinkedList()
	return queue
}

// since the queue.item.topSentinel is accessed quite often
// define a function to access it
func (queue *Queue) top() *Cell {
	return queue.items.topSentinel
}

// since the queue.item.bottomSentinel is accessed quite often
// define a function to access it
func (queue *Queue) bottom() *Cell {
	return queue.items.bottomSentinel
}

func (queue *Queue) enqueue(item string) {
	newCell := &Cell{data: item}
	queue.bottom().addBefore(newCell)
}

func (queue *Queue) dequeue() string {

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

func (queue *Queue) isEmpty() bool {
	// queue is empty if all there is are the top and bottom sentinels
	// and they are pointing to each other
	top := queue.top()
	bottom := queue.bottom()

	return top.next == bottom && bottom.prev == top
}
func (queue *Queue) toString(sep string) string {
	return queue.items.toString(sep)
}
