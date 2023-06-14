package main

// Returns true iff the linked list contains a loop
func (list *LinkedList) hasLoop() bool {
	var fast, slow *Cell

	// initialise both pointers to point 
	// to the sentinel, w/c is guaranteed to be not nil
	fast = list.sentinel
	slow = list.sentinel

	for {
		// when either fast or slow pointer has reached a nil:
		// return false - i.e. no loop
		// otherwise advance the pointers
		// fast takes two steps at a time
		// slow takes one step at a time
		if fast.next == nil {
			return false
		} else {
			if fast.next.next == nil {
				return false
			} else {
				fast = fast.next.next
			}
		}
		
		if slow.next == nil {
			return false
		} else {
			slow = slow.next 
		}
		// when slow finally catches up with fast
		// return true - i.e. there is a loop
		if slow == fast {
			return true
		}
	}
}