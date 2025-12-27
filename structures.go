package main

import (
	"fmt"
)


func main(){

	type xyz struct{
		x int
		y int
		z int 
	}

	var s1 xyz
	fmt.Println(s1.y,s1.z)

	p1 := xyz{23,12,-2}
	p2 := xyz{z:12,y:13}

	fmt.Println(p1)
	fmt.Println(p2)


	pSlice := [4]xyz{}
	pSlice[2] = p1
	pSlice[0] = p2

	fmt.Println(pSlice)
	p2 = xyz{1,2,3}
	fmt.Println(pSlice)

}