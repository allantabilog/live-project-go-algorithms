package main

type Queue struct {
	items DoublyLinkedList
}

func makeQueue() Queue {
	var queue Queue
	queue.items = makeDoublyLinkedList()
	return queue
}

func (queue *Queue) enqueue(item string) {
	var lastCell = queue.items.bottomSentinel.prev
	var newCell *Cell = &Cell{data: item}
	lastCell.next = newCell
	newCell.prev = lastCell
	newCell.next = queue.items.bottomSentinel
	queue.items.bottomSentinel.prev = newCell
}

func (queue *Queue) dequeue() string {

	// if the queue is empty, calling dequeue() will panic
	if queue.isEmpty() {
		panic("Dequeueing an empty queue")
	}

	var cellToDequeue = queue.items.topSentinel.next
	var newFirstCell = cellToDequeue.next

	queue.items.topSentinel.next = newFirstCell
	newFirstCell.prev = queue.items.topSentinel

	return cellToDequeue.data
}

func (queue *Queue) isEmpty() bool {
	// queue is empty if all there is are the top and bottom sentinels
	// and they are pointing to each other
	top := queue.items.topSentinel
	bottom := queue.items.bottomSentinel

	return top.next == bottom && bottom.prev == top
}
func (queue *Queue) toString(sep string) string {
	return queue.items.toString(sep)
}
