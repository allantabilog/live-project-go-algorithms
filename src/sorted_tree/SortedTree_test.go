package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertSorted(t *testing.T) {
	var tree SortedTree 
	tree.setRoot("")
	var root = tree.root
	root.insertValue("I")
    root.insertValue("G")
    root.insertValue("C")
    root.insertValue("E")
    root.insertValue("B")
    root.insertValue("K")
    root.insertValue("S")
    root.insertValue("Q")
    root.insertValue("M")
    root.insertValue("F")

	var sortedValues = root.right.inorder()
	assert.Equal(t, "B C E F G I K M Q S", sortedValues)
}

func TestFind_FindsExisting(t *testing.T) {
	var tree SortedTree 
	tree.setRoot("")
	var root = tree.root
	root.insertValue("I")
    root.insertValue("G")
    root.insertValue("C")
    root.insertValue("E")
    root.insertValue("B")
    root.insertValue("K")
    root.insertValue("S")
    root.insertValue("Q")
    root.insertValue("M")
    root.insertValue("F")

	var searchNode *Node
	searchNode = root.findValue("C")
	assert.Equal(t, "C", searchNode.data)
	
	searchNode = root.findValue("Q")
	assert.Equal(t, "Q", searchNode.data)
}

func TestFind_ReturnsNilForNonExisting(t *testing.T) {
	var tree SortedTree 
	tree.setRoot("")
	var root = tree.root
	root.insertValue("I")
    root.insertValue("G")
    root.insertValue("C")
    root.insertValue("E")
    root.insertValue("B")
    root.insertValue("K")
    root.insertValue("S")
    root.insertValue("Q")
    root.insertValue("M")
    root.insertValue("F")

	var searchNode *Node
	searchNode = root.findValue("J")
	assert.Nil(t, searchNode)

	searchNode = root.findValue("Z")
	assert.Nil(t, searchNode)
}