package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringQueue_Enqueue(t *testing.T) {
	queue := makeQueue[string]("top", "bottom")
	queue.enqueue("1")
	queue.enqueue("2")
	queue.enqueue("3")

	assert.Equal(t,
		"top <-> 1 <-> 2 <-> 3 <-> bottom",
		queue.toString("<->"),
	)
}

func TestStringDequeue_Empty(t *testing.T) {
	queue := makeQueue[string]("top", "bottom")
	assert.Panics(t, func() {
		queue.dequeue()
	})
}

func TestStringDequeue_NonEmpty(t *testing.T) {
	queue := makeQueue[string]("top", "bottom")
	queue.enqueue("1")
	queue.enqueue("2")
	queue.enqueue("3")

	assert.Equal(t, "1", queue.dequeue())
	fmt.Println(queue.toString(">"))
	assert.Equal(t, "2", queue.dequeue())
	fmt.Println(queue.toString(">"))
	assert.Equal(t, "3", queue.dequeue())
	assert.True(t, queue.isEmpty())
}

type Node struct {
	data string
	next *Node
}

func TestNodeQueue_Enqueue(t *testing.T) {
	var top Node = Node{"top", nil}
	var bottom Node = Node{"bottom", nil}

	queue := makeQueue[Node](top, bottom)
	queue.enqueue(Node{data: "data", next: nil})

	assert.Equal(t,
		"{top <nil>} > {data <nil>} > {bottom <nil>}",
		queue.toString(">"),
	)

}
