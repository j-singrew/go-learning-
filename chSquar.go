package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var times int 

func f1(cc chan chan int,f chan bool){
	c := make(chan int)
	cc <- c
	defer close(c)

	sum := 0
	select {
	case x := <- c:
		for i := 0;i < x;i++{
			sum = sum + 1
		}
		c <- sum
	case <-f:
		return 
	}
}

func main(){
	argument := os.Args
	if len(argument) != 2{
		fmt.Println("Need just one integer argument!")
		return
	}

	times ,err := strconv.Atoi(argument[1])
	if err != nil{
		fmt.Println(err)
		return 
	}

	cc := make(chan chan int)

for i := 1;i < times + 1;i++{
	f := make(chan bool)
	go f1(cc,f)
	ch := <-cc
	ch <- i
	for sum := range ch{
		fmt.Print("Sum(", i, ")=", sum)
	}
	fmt.Println()
	time.Sleep(time.Second)
	close(f)
}
}