package main

import (
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
