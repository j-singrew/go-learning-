package main

import(
	"fmt"
)

type a struct {
	XX int 
	YY int
}

type b struct{
	AA string
	XX int
}

type c srtuct{
	A a 
	B b	
}

func (A a) A(){
	fmt.Println("Function A() for A")
}

func (B b) A(){
	fmt.Println("Function A() for B")
}


func main(){
	Var i c
	i.A.A()
	i.B.A()
}