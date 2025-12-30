package main

import (
	"fmt"
	"time"
)


func writeToChannel(c chan int, x int){
	fmt.Println(x)
	c <- x
	close(c)
	fmt.Println(x)
}

func main(){
	time.Sleep(1 * time.Second)
}