package main

import (
	"fmt"
	"testing"
)

func TestInsertValue(t *testing.T) {
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

	fmt.Printf("Sorted values: %s", root.right.inorder())

	var searchNode *Node
	searchNode = root.findValue("C")
	showSearchResult(searchNode)
	searchNode = root.findValue("Q")
	showSearchResult(searchNode)
	searchNode = root.findValue("J")
	showSearchResult(searchNode)
	
}

func showSearchResult(node *Node) {
	if node != nil {
		fmt.Println("Node found. data: ", node.data)
	} else {
		fmt.Println("Node not found ")
	}
}