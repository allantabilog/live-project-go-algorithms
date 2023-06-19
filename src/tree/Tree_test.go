package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTree(t *testing.T) {
	assert.True(t, true)
	var tree = buildTree()
	fmt.Println(tree.root.displayIndented(" ", 0))
}
