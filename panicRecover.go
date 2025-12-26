package main
import (
	"fmt"
)

func a(){
	fmt.Println("inside a()")
	defer func(){
		if c := recover();c != nil{
			fmt.Println("Recvover inside a()!")
		}
	}()
	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited!")
	fmt.Println("Exiteing a()")
}

func b(){
	fmt.Println("Inside b()")
	panic("panic in b()!")
	fmt.Println("Exiting b()")
}


func main(){
	a()
	fmt.Println("main() ended!")
}
