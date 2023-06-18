package main

type Deque struct {
	items DoublyLinkedList
}

func makeDeque() Deque {
	var deque Deque
	deque.items = makeDoublyLinkedList()
	return deque
}

func (deque *Deque) top() *Cell {
	return deque.items.topSentinel
}

func (deque *Deque) bottom() *Cell {
	return deque.items.bottomSentinel
}

func (deque *Deque) isEmpty() bool {
	top := deque.top()
	bottom := deque.bottom()
	return top.next == bottom && bottom.prev == top
}

func (deque *Deque) pushTop(item string) {
	deque.top().addAfter(&Cell{data: item})
}

func (deque *Deque) pushBottom(item string) {
	deque.bottom().addBefore(&Cell{data: item})
}

func (deque *Deque) popTop() string {
	if deque.isEmpty() {
		panic("Popping an empty deque")
	}
	top := deque.top()
	// save the popped data before removing the cell
	popped := top.next

	// re-point the next and prev links
	top.next = top.next.next
	top.next.prev = popped.prev

	// reset the prev and next links of popped cell
	popped.prev = nil
	popped.next = nil

	return popped.data
}

func (deque *Deque) popBottom() string {
	if deque.isEmpty() {
		panic("Popping an empty deque")
	}
	bottom := deque.bottom()
	popped := bottom.prev

	// re-point the next and prev links
	bottom.prev = bottom.prev.prev
	bottom.prev.next = popped.next
	return popped.data
}

func (deque *Deque) toString(sep string) string {
	return deque.items.toString(sep)
}

// use this for printing the structure
// in the case when there are loops
func (deque *Deque) toStringMax(sep string, max int) string {
	return deque.items.toStringMax(sep, max)
}
