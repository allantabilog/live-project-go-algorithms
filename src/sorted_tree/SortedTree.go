package main

type Node struct {
	data        string
	left, right *Node
}

type SortedTree struct {
	root *Node
}

func (tree *SortedTree) setRoot(data string) {
        tree.root = &Node{data: data}
}

func (node *Node) setLeft(data string) {
        node.left = &Node{data: data}
}

func (node *Node) setRight(data string) {
        node.right = &Node{data: data}
}

func (node *Node) insertValue(newValue string) {
	if node.data >= newValue {
		if node.left != nil {
			node.left.insertValue(newValue)
		} else {
			node.setLeft(newValue)
		}
	} else {
		if node.right != nil {
			node.right.insertValue(newValue)
		} else {
			node.setRight(newValue)
		}
	}
}

func (node *Node) inorder() string {
	var result string

	if node.left != nil {
			result += node.left.inorder() + " "
	}

	result += node.data

	if node.right != nil {
			result += " " + node.right.inorder()
	}

	return result
}

func (node *Node) findValue(value string) *Node {
	if node.data == value {
		return node
	} else if node.data >= value {
		if node.left != nil {
			return node.left.findValue(value)
		} else {
			return nil
		}
	} else {
		if node.right != nil {
			return node.right.findValue(value)
		} else {
			return nil
		}
	}
}
