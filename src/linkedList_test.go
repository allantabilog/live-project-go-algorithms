package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeLinkedList(t *testing.T){
	list := makeLinkedList()
	assert.NotNil(t, list)
	assert.NotNil(t, list.sentinel)
	assert.Equal(t, list.sentinel.data, "Sentinel")
	assert.Nil(t, list.sentinel.next)
}

func TestAddAfter_simple(t *testing.T) {
	var first = Cell{data: "first"}
	var second = Cell{data: "second"}
	first.addAfter(&second)

	assert.Equal(t, first.next, &second)
	assert.Nil(t, second.next)
}
func TestAddAfter_insert(t *testing.T) {
	var first = Cell{data: "first"}
	var second = Cell{data: "second"}
	var insert = Cell{data: "insert"}

	// setup so that first -> second
	first.addAfter(&second)

	assert.Equal(t, first.next, &second)
	assert.Nil(t, second.next)

    // Now add the cell "insert" after "first"
	// expected first -> insert -> second
	first.addAfter(&insert)
	assert.Equal(t, first.next, &insert)
	assert.Equal(t, insert.next, &second)
	assert.Nil(t, second.next)
}

func TestDeleteAfter_simple(t *testing.T){
	var first = Cell{data: "first"}
	var second = Cell{data: "second"}
	// setup so that first -> second
	first.addAfter(&second)

	var deleted = first.deleteAfter()
	assert.Equal(t, deleted, &second)
	assert.Nil(t, first.next)
}

func TestDeleteAfter_nil(t *testing.T){
	assert.Panics(t, func() {
		var first = Cell{data: "first"}
		first.deleteAfter()
	})
}