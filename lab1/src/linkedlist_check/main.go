package main

import (
	"github.com/KevinZonda/The-Go-Book/lab1/src/linkedlist"
	"github.com/KevinZonda/The-Go-Book/lab1/src/node"
	"github.com/KevinZonda/The-Go-Book/tester"
)

func main() {
	tester.RunTestNodes([]tester.TestNode{
		{Name: "NewLinkedList", Func: newLinkedList},
		{Name: "HasValue", Func: has},
		{Name: "AddNode", Func: add},
		{Name: "RemoveNode", Func: remove},
	})
}

func newLinkedList() bool {
	ll := linkedlist.NewLinkedList(nil)
	if ll == nil {
		return tester.F("LinkedList shouldn't be nil")
	}
	if ll.Head() != nil {
		return tester.F("LinkedList Head should be nil")
	}
	n := node.NewLinkedListNode("Hello")
	ll = linkedlist.NewLinkedList(n)

	if ll == nil {
		return tester.F("LinkedList shouldnt be nil")
	}
	if ll.Head() == nil {
		return tester.F("LinkedList Head should be nil")
	}

	if ll.Head().Value() != "Hello" {
		return tester.F("Wrong HEAD")
	}
	if ll.Head().Next() != nil {
		return tester.F("Wrong NEXT")
	}

	n.SetNext(node.NewLinkedListNode("World"))
	if ll.Head().Next() == nil || ll.Head().Next().Next() != nil {
		return tester.F("Wrong NEXT")
	}
	if ll.Head().Next().Value() != "World" {
		return tester.F("Not World at second")
	}
	return true

}

func has() bool {
	ll := linkedlist.NewLinkedList(nil)
	if ll.Has("666") {
		return tester.F("NOT AVAIL 666")
	}
	n := node.NewLinkedListNode("Hello")
	ll = linkedlist.NewLinkedList(n)
	if !ll.Has("Hello") {
		return tester.F("SHOULD BE HELLO")
	}
	if ll.Has("666") {
		return tester.F("NOT AVAIL 666 #2")
	}
	return true
}

func add() bool {
	ll := linkedlist.NewLinkedList(nil)
	ll.Add("World")
	ll.Add("Hello")
	if !(ll.Has("Hello") && ll.Has("World")) {
		return tester.F("Not in side")
	}
	if ll.Head().Length() != 2 {
		return tester.F("Not correct length")
	}
	return true
}

func remove() bool {
	ll := linkedlist.NewLinkedList(nil)
	ll.Add("World")
	ll.Add("Hello")

	ll.Remove("Hello")
	if ll.Has("Hello") {
		return tester.F("Hello is inside")
	}
	if ll.Head().Length() != 1 {
		return tester.F("Leng != 1")
	}

	ll.Remove("World")
	if ll.Has("World") {
		return tester.F("World is inside")
	}
	if ll.Head().Length() != 0 {
		return tester.F("Leng != 0")
	}
	return true
}
