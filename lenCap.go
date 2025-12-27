package main

import (
	"fmt"
)

func printSlice(x []int){
	for _,number := range x {
		fmt.Print(number," ")
	}
	fmt.Println()
}


func main(){
	aSlice := []int{-1,0,4}
	fmt.Printf("aslice: ")
	printSlice(aSlice)

	fmt.Printf("cap: %d,length : %d\n",cap(aSlice),len(aSlice))
	aSlice = append(aSlice,-110)
	fmt.Printf("aslice: ")
	printSlice(aSlice)
	fmt.Printf("cap: %d,length: %d\n",cap(aSlice),len(aSlice))
}