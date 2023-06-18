package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	deque := makeDeque()
	assert.True(t, deque.isEmpty())
}

func TestPushTop(t *testing.T) {
	deque := makeDeque()

	deque.pushTop("1")
	deque.pushTop("2")
	deque.pushTop("3")

	assert.Equal(t,
		"top_sentinel > 3 > 2 > 1 > bottom_sentinel",
		deque.toString(">"),
	)
}

func TestPushBottom(t *testing.T) {
	deque := makeDeque()

	deque.pushBottom("1")
	deque.pushBottom("2")
	deque.pushBottom("3")

	assert.Equal(t,
		"top_sentinel > 1 > 2 > 3 > bottom_sentinel",
		deque.toString(">"),
	)
}

func TestPopTop_Empty(t *testing.T) {
	deque := makeDeque()
	assert.Panics(t, func() {
		deque.popTop()
	})
}

func TestPopTop_AfterPushTop(t *testing.T) {
	deque := makeDeque()
	deque.pushTop("1")
	deque.pushTop("2")
	deque.pushTop("3")

	assert.Equal(t, "3", deque.popTop())
	assert.Equal(t, "2", deque.popTop())
	assert.Equal(t, "1", deque.popTop())
	assert.True(t, deque.isEmpty())
}

func TestPopBottom_AfterPushBottom(t *testing.T) {
	deque := makeDeque()
	deque.pushBottom("1")
	deque.pushBottom("2")
	deque.pushBottom("3")
	assert.Equal(t, "3", deque.popBottom())
	assert.Equal(t, "2", deque.popBottom())
	assert.Equal(t, "1", deque.popBottom())
	assert.True(t, deque.isEmpty())
}

func TestPopTop_AfterPushBottom(t *testing.T) {
	deque := makeDeque()
	deque.pushBottom("1")
	deque.pushBottom("2")
	deque.pushBottom("3")
	assert.Equal(t, "1", deque.popTop())
	assert.Equal(t, "2", deque.popTop())
	assert.Equal(t, "3", deque.popTop())
	assert.True(t, deque.isEmpty())
}

func TestPopBottom_AfterPushTop(t *testing.T) {
	deque := makeDeque()
	deque.pushTop("1")
	deque.pushTop("2")
	deque.pushTop("3")
	assert.Equal(t, "1", deque.popBottom())
	assert.Equal(t, "2", deque.popBottom())
	assert.Equal(t, "3", deque.popBottom())
	assert.True(t, deque.isEmpty())
}
