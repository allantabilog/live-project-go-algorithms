package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTree(t *testing.T) {
	var tree = buildTree()
	var expectedOutput = `A
 B
  D
  E
   G
 C
  F
   H
    I
    J
`
	assert.Equal(t, expectedOutput, tree.root.displayIndented(" ", 0))
}

func TestPreorder(t *testing.T) {
	var tree = buildTree()
	var expectedTraversal = "A B D E G C F H I J"
	assert.Equal(t, expectedTraversal, tree.root.preorder())
}

func TestInorder(t *testing.T) {
	var tree = buildTree()
	var expectedTraversal = "D B G E A C I H J F"
	assert.Equal(t, expectedTraversal, tree.root.inorder())
}

func TestPostorder(t *testing.T) {
	var tree = buildTree()
	var expectedTraversal = "D G E B I J H F C A"
	assert.Equal(t, expectedTraversal, tree.root.postorder())
}
