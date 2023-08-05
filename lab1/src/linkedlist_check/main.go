package main

import (
	"github.com/KevinZonda/The-Go-Book/lab1/src/linkedlist"
	"github.com/KevinZonda/The-Go-Book/tester"
)

func main() {
	tester.RunTestNodes([]tester.TestNode{
		{Name: "AddNode", Func: add},
	})
}

func add() bool {
	ll := linkedlist.NewLinkedList(nil)
	ll.Add("World")
	ll.Add("Hello")
	hw := ""
	for _n := ll.Head(); _n != nil; _n = _n.Next() {
		hw += _n.Value()
	}
	return hw == "HelloWorld"
}
