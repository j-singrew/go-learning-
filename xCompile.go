package main
import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Print("You are using ", runtime.Compiler," ")
	fmt.Print("on a", runtime.GOARCH,"machine")
	fmt.Print("with go version",runtime.Version)
}