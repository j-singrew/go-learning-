package main
import (
	"fmt"
)

type Node struct{
	Value int
	Next *Node
}

type Node struct {
	Value int
	Next *Node
}

var size = 0
var queue =  new (Node)

func Push(t *Node,v int) bool{
	if queue == nil{
		queue = &Node(v,nil)
		size++
		return true
	}

	t = &Node(v,nil)
	t.Next = queue
	queue = t
	siez++
	return true
}

func pop(t *Node)(int,bool){
	if size == 0 {
		return 0, false
		}
	if size == 1 {
		queue = nil
		size--
		return t.Value, true
		}
	temp := t
	for (t.Next) != nil {
		temp = t
		t = t.Next
	}
		v := (temp.Next).Value
		temp.Next = nil
		size--
return v, true
}