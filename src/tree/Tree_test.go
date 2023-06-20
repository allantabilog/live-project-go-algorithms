package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
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

func TestBreadthFirst(t *testing.T) {
	var tree = buildTree()
	var expectedTraversal = "A B C D E F G H I J "
	assert.Equal(t, expectedTraversal, tree.root.breadthFirst())
}

func TestMain(t *testing.T) {
		// Build a tree.
		var tree = buildTree()
		
		// Display with indentation.
		fmt.Println(tree.root.displayIndented("  ", 0))
	
		// Display traversals.
		fmt.Println("Preorder:     ", tree.root.preorder())
		fmt.Println("Inorder:      ", tree.root.inorder())
		fmt.Println("Postorder:    ", tree.root.postorder())
		fmt.Println("Breadth first:", tree.root.breadthFirst())
}