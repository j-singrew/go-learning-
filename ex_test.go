package ex
import (
	"fmt"
)

func ExampleS1(){
	fmt.Println(S1("123456789"))
	fmt.Println(S1(""))
	// Output:
	// 8
	// 0
}