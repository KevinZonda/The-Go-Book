package main

import (
	"github.com/KevinZonda/The-Go-Book/lab1/src/node"
	"github.com/KevinZonda/The-Go-Book/tester"
)

func main() {
	tester.RunTestNodes([]tester.TestNode{
		{Name: "NewNode", Func: newNode},
		{Name: "SetNext", Func: setNext},
	})
}

func newNode() bool {
	n := node.NewLinkedListNode("666")
	if n.Value() != "666" {
		return false
	}
	if n.Next() != nil {
		return false
	}
	return true
}

func setNext() bool {
	head := node.NewLinkedListNode("Hello")
	head.SetNext(node.NewLinkedListNode("World"))
	hw := ""
	for _n := head; _n != nil; _n = _n.Next() {
		hw += _n.Value()
	}
	return hw == "HelloWorld"
}
