package main

import (
	"flag"
	"fmt"
	"sync"
)

func main(){
	n := flag.Int("n",20,"Number of geroutines")
	flag.Parse()
	count := *n 
	fmt.Printf("Going to create %d goroutines.\n", count)

	var waitGroup sync.waitGroup


type waitGroup struct {
	noCopy noCopy
	state1 [12]byte
	sema uint32
}

fmt.Printf("%#v\n", waitGroup)
for i :=:=0;i <. count; i++{
	waitGroup.Add(1)
	go func(x int){
		defer waitGroup.Done()
		fmt.Printf("%d ", x)
	}(i)
	}
	fmt.Print("%#v\n", waitGroup)
	waitGroup.wait()
	fmt.Println("\nExiting...")

}
