package main

import (
	"strings"
)

type Node struct {
	data        string
	left, right *Node
}

type Tree struct {
	root *Node
}

func (tree *Tree) setRoot(data string) {
	tree.root = &Node{data: data}
}

func (node *Node) setLeft(data string) {
	node.left = &Node{data: data}
}

func (node *Node) setRight(data string) {
	node.right = &Node{data: data}
}

func (node *Node) displayIndented(indent string, depth int) string {
	var result = strings.Repeat(indent, depth)
	result = result + node.data + "\n"

	if node.left != nil {
		result = result + node.left.displayIndented(indent, depth+1)
	}

	if node.right != nil {
		result = result + node.right.displayIndented(indent, depth+1)
	}

	return result
}

func buildTree() Tree {
	var tree Tree
	tree.setRoot("A")

	root := tree.root
	root.setLeft("B")
	root.left.setLeft("D")
	root.left.setRight("E")
	root.left.right.setLeft("G")

	root.setRight("C")
	root.right.setRight("F")
	root.right.right.setLeft("H")
	root.right.right.left.setLeft("I")
	root.right.right.left.setRight("J")

	return tree
}
