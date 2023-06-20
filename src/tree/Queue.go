package main

type Cell struct {
	data *Node
	prev *Cell
	next *Cell
}

type DoublyLinkedListNode struct {
	topSentinel    *Cell
	bottomSentinel *Cell
}

func makeDoublyLinkedListNode() DoublyLinkedListNode {
	list := DoublyLinkedListNode {
		topSentinel:    &Cell{data: nil},
		bottomSentinel: &Cell{data: nil},
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


type QueueNode struct {
	items DoublyLinkedListNode
}

func makeQueueNode() QueueNode {
	var queue QueueNode
	queue.items = makeDoublyLinkedListNode()
	return queue
}

// since the queue.item.topSentinel is accessed quite often
// define a function to access it
func (queue *QueueNode) top() *Cell {
	return queue.items.topSentinel
}

// since the queue.item.bottomSentinel is accessed quite often
// define a function to access it
func (queue *QueueNode) bottom() *Cell {
	return queue.items.bottomSentinel
}

func (queue *QueueNode) enqueue(item *Node) {
	newCell := &Cell{data: item}
	queue.bottom().addBefore(newCell)
}

func (queue *QueueNode) dequeue() *Node {

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

func (queue *QueueNode) isEmpty() bool {
	// queue is empty if all there is are the top and bottom sentinels
	// and they are pointing to each other
	top := queue.top()
	bottom := queue.bottom()

	return top.next == bottom && bottom.prev == top
}
