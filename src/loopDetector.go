package main

// Returns true iff the linked list contains a loop
func (list *LinkedList) hasLoop() bool {
	var fast, slow *Cell

	fast = list.sentinel
	slow = list.sentinel

	for {
		fast = fast.next.next 
		slow = slow.next

		if fast == nil {
			return false
		}
		if slow == fast {
			return true
		}
	}
}