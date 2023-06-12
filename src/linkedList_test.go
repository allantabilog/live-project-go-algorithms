package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeLinkedList(t *testing.T){
	list := makeLinkedList()
	assert.NotNil(t, list)
	assert.NotNil(t, list.sentinel)
	assert.Equal(t, "Sentinel", list.sentinel.data)
	assert.Nil(t, list.sentinel.next)
}

func TestAddAfter_simple(t *testing.T) {
	var first = Cell{data: "first"}
	var second = Cell{data: "second"}
	first.addAfter(&second)

	assert.Equal(t, &second, first.next)
	assert.Nil(t, second.next)
}
func TestAddAfter_insert(t *testing.T) {
	var first = Cell{data: "first"}
	var second = Cell{data: "second"}
	var insert = Cell{data: "insert"}

	// setup so that first -> second
	first.addAfter(&second)

	assert.Equal(t, &second, first.next)
	assert.Nil(t, second.next)

    // Now add the cell "insert" after "first"
	// expected first -> insert -> second
	first.addAfter(&insert)
	assert.Equal(t, &insert, first.next)
	assert.Equal(t, &second, insert.next)
	assert.Nil(t, second.next)
}

func TestBuildList(t *testing.T){
	var list = makeLinkedList()
	var sentinel = list.sentinel
	var first = Cell{data: "first"}
	var second = Cell{data: "second"}
	var third = Cell{data: "third"}

	sentinel.addAfter(&first)
	first.addAfter(&second)
	second.addAfter(&third)

	// expect: sentinel -> first -> second -> third
	assert.Equal(t, &first, sentinel.next )
	assert.Equal(t, &second, first.next  )
	assert.Equal(t, &third, second.next  )

}

func TestDeleteAfter_simple(t *testing.T){
	var first = Cell{data: "first"}
	var second = Cell{data: "second"}
	// setup so that first -> second
	first.addAfter(&second)

	var deleted = first.deleteAfter()
	assert.Equal(t, &second, deleted)
	assert.Nil(t, first.next)
}

func TestDeleteAfter_nil(t *testing.T){
	assert.Panics(t, func() {
		var first = Cell{data: "first"}
		first.deleteAfter()
	})
}

func TestAddRange_Empty(t *testing.T) {
	// initially empty list
	var list = makeLinkedList()
	list.addRange([]string{"1", "2", "3"})

	// expect sentinel -> 1 -> 2 -> 3
	assert.Equal(t, "1", list.sentinel.next.data)
	assert.Equal(t, "2", list.sentinel.next.next.data)
	assert.Equal(t, "3", list.sentinel.next.next.next.data)
}

func TestAddRange_NonEmpty(t *testing.T) {
	// initially nonempty list
	var list = makeLinkedList()
	list.sentinel.addAfter(&Cell{data: "0"})
	list.addRange([]string{"1", "2", "3"})

	// expect sentinel -> 1 -> 2 -> 3
	assert.Equal(t, "0", list.sentinel.next.data)
	assert.Equal(t, "1", list.sentinel.next.next.data)
	assert.Equal(t, "2", list.sentinel.next.next.next.data)
	assert.Equal(t, "3", list.sentinel.next.next.next.next.data)
}

func TestToString(t *testing.T) {
	var list = makeLinkedList()
	assert.Equal(t, list.toString("->"), "Sentinel")

	list.addRange([]string{"1", "2", "3", "4", "5"})
	assert.Equal(t, "Sentinel > 1 > 2 > 3 > 4 > 5", list.toString(" > "))
}