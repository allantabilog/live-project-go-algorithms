package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeDoublyLinkedList(t *testing.T) {
	list := makeDoublyLinkedList()
	assert.True(t, list.topSentinel.next == list.bottomSentinel)
	assert.True(t, list.bottomSentinel.prev == list.topSentinel)
	assert.Nil(t, list.topSentinel.prev)
	assert.Nil(t, list.bottomSentinel.next)

}
func TestAddAfter_Empty(t *testing.T) {
	list := makeDoublyLinkedList()
	newCell := Cell{data: "first"}
	list.topSentinel.addAfter(&newCell)

	assert.Equal(t, &newCell, list.topSentinel.next)
	assert.Equal(t, &newCell, list.bottomSentinel.prev)
}

func TestAddAfter_NonEmpty(t *testing.T) {
	list := makeDoublyLinkedList()
	firstCell := Cell{data: "first"}
	secondCell := Cell{data: "second"}
	thirdCell := Cell{data: "third"}

	list.topSentinel.addAfter(&firstCell)
	list.topSentinel.next.addAfter(&secondCell)

	assert.Equal(t,
		"top_sentinel <-> first <-> second <-> bottom_sentinel",
		list.toString("<->"),
	)

	list.topSentinel.next.addAfter(&thirdCell)
	assert.Equal(t,
		"top_sentinel <-> first <-> third <-> second <-> bottom_sentinel",
		list.toString("<->"),
	)

}

func TestAddBefore_Empty(t *testing.T) {
	list := makeDoublyLinkedList()
	newCell := Cell{data: "first"}
	list.bottomSentinel.addBefore(&newCell)

	assert.Equal(t, &newCell, list.topSentinel.next)
	assert.Equal(t, &newCell, list.bottomSentinel.prev)
}

func TestAddBefore_NonEmpty(t *testing.T) {
	list := makeDoublyLinkedList()
	firstCell := Cell{data: "first"}
	secondCell := Cell{data: "second"}
	thirdCell := Cell{data: "third"}

	list.topSentinel.addAfter(&firstCell)
	list.topSentinel.next.addAfter(&secondCell)

	assert.Equal(t,
		"top_sentinel <-> first <-> second <-> bottom_sentinel",
		list.toString("<->"),
	)

	list.topSentinel.next.next.addBefore(&thirdCell)
	assert.Equal(t,
		"top_sentinel <-> first <-> third <-> second <-> bottom_sentinel",
		list.toString("<->"),
	)

}
func TestAddRange(t *testing.T) {
	list := makeDoublyLinkedList()
	list.addRange([]string{"1", "2", "3"})
	assert.Equal(t,
		"top_sentinel <-> 1 <-> 2 <-> 3 <-> bottom_sentinel",
		list.toString("<->"),
	)
}

// test following all the links forward from the topSentinel
// via .next links
func TestForwardTraverse(t *testing.T) {
	list := makeDoublyLinkedList()
	list.addRange([]string{"1", "2", "3"})
	// test following all the links forward from the topSentinel
	var b strings.Builder
	var sep string = ":"

	for curr := list.topSentinel; curr != nil; curr = curr.next {
		if curr.next != nil {
			fmt.Fprintf(&b, "%s %s ", curr.data, sep)
		} else {
			fmt.Fprintf(&b, "%s", curr.data)
		}
	}
	assert.Equal(t, "top_sentinel : 1 : 2 : 3 : bottom_sentinel", b.String())
}

// test following all the links backward from the bottomSentinel
// via .prev links
func TestBackwardTraverse(t *testing.T) {
	list := makeDoublyLinkedList()
	list.addRange([]string{"1", "2", "3"})
	// test following all the links forward from the topSentinel
	var b strings.Builder
	var sep string = ":"

	for curr := list.bottomSentinel; curr != nil; curr = curr.prev {
		if curr.prev != nil {
			fmt.Fprintf(&b, "%s %s ", curr.data, sep)
		} else {
			fmt.Fprintf(&b, "%s", curr.data)
		}
	}
	assert.Equal(t, "bottom_sentinel : 3 : 2 : 1 : top_sentinel", b.String())
}

func TestToString(t *testing.T) {
	list := makeDoublyLinkedList()
	firstCell := Cell{data: "first"}
	secondCell := Cell{data: "second"}
	list.topSentinel.addAfter(&firstCell)
	list.topSentinel.next.addAfter(&secondCell)

	assert.Equal(t,
		"top_sentinel <-> first <-> second <-> bottom_sentinel",
		list.toString("<->"),
	)
}
