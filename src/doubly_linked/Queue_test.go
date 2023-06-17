package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	queue := makeQueue()
	queue.enqueue("1")
	queue.enqueue("2")
	queue.enqueue("3")

	assert.Equal(t,
		"top_sentinel <-> 1 <-> 2 <-> 3 <-> bottom_sentinel",
		queue.toString("<->"),
	)
}

func TestDequeue_Empty(t *testing.T) {
	queue := makeQueue()
	assert.Panics(t, func() {
		queue.dequeue()
	})
}

func TestDequeue_NonEmpty(t *testing.T) {
	queue := makeQueue()
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
