package main
import (
	"fmt"
)

type Node struct {
	Value int
	Previous *Node
	Next *Node
}

func addNode(t *Node,v in) int{
	if root == nil{
		fmt.Println("->Empty list!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}
func  reverse(t *Node){
	if  t == nil{
		fmt.Printf("-> Empty list!")
		return 
	}
	temp := t
	for t != nil{
		temp = t
		t = t.Next
	}

	for temp.Previous != nil{
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Previous
	}
	fmt.Printf("%d -> ", temp.Value)
	fmt.Println()
}

func size(t *Node) int {
	if t == nil{
		fmt.Println("-> Empty list!")
		return 0
	}
	n := 0
	for t != nil{
		n++
		t = t.Next
		
	}
	return n
}

func lookupNode(t *Node, v int)bool {
	if root == nil{
		return false
	}
	if v == t.value{
		return true
	}
	if t.Next == nil{
		return false
	}
	return lookupNode(t.Next,v)

}


var  root = new(Node)

func main(){
	fmt.Printf(root)
	root = nil
	traverse(root)
	addNode(root, 1)
	addNode(root, 1)
	traverse(root)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 0)
	addNode(root, 0)
	traverse(root)

	addNode(root, 100)
	fmt.Println("Size:", size(root))
	traverse(root)
	reverse(root)
}