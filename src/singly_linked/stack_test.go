package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeStack(t *testing.T) {
	stack := makeStack()
	assert.NotNil(t, stack)
	assert.True(t, stack.isEmpty())
}

func TestPush(t *testing.T) {
	stack := makeStack()
	assert.True(t, stack.isEmpty())

	stack.push("Item 1")
	assert.False(t, stack.isEmpty())
	assert.Equal(t, "Item 1", stack.items.sentinel.next.data)

	stack.push("Item 2")
	assert.Equal(t, "Item 2", stack.items.sentinel.next.data)

	stack.push("Item 3")
	assert.Equal(t, "Item 3", stack.items.sentinel.next.data)

	assert.Equal(t, 3, stack.length())
}

func TestPop_Empty(t *testing.T) {
	stack := makeStack()
	assert.True(t, stack.isEmpty())

	assert.Panics(t, func() {
		stack.pop()
	})

	stack.push("Item 1")

	poppedItem := stack.pop()
	assert.Equal(t, "Item 1", poppedItem)
	assert.True(t, stack.isEmpty())
}

func TestPop_NonEmpty(t *testing.T) {
	stack := makeStack()
	stack.push("Item 1")

	poppedItem := stack.pop()
	assert.Equal(t, "Item 1", poppedItem)
	assert.True(t, stack.isEmpty())
}

func TestLIFO(t *testing.T) {
	stack := makeStack()

	stack.push("1")
	stack.push("2")
	stack.push("3")
	assert.Equal(t, "3", stack.pop())
	assert.Equal(t, "2", stack.pop())
	assert.Equal(t, "1", stack.pop())
	assert.True(t, stack.isEmpty())
}

func TestLIFO_ShowContents(t *testing.T) {
	stack := makeStack()

	stack.push("Apple")
	stack.push("Banana")
	stack.push("Coconut")
	stack.push("Date")

	for !stack.isEmpty() {
		fmt.Printf("Popped %-7s Remaining %d: %v\n",
			stack.pop(),
			stack.length(),
			stack.toArray())
	}

}
