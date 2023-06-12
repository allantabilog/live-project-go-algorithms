package main

import "fmt"

type Cell struct {
	data string
	next *Cell
}

type LinkedList struct {
	sentinel *Cell
}
func main() {

	listOne := makeLinkedList()
	fmt.Println(listOne);

}	
func makeLinkedList() LinkedList{
	// return LinkedList{ sentinel: nil } // my original code
	return LinkedList {
		sentinel: &Cell{data: "Sentinel", next: nil},
	}
}

// Add the cell <cell> after the cell <me>
func (me *Cell) addAfter(cell *Cell) {
	cell.next = me.next
	me.next = cell
}
func (me *Cell) deleteAfter() *Cell {
	deletedCell := me.next 
	if deletedCell == nil {
		panic("No cell to delete")
	}
	me.next = nil 
	return deletedCell
}