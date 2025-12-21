package main

import (
	"fmt"
)

func main(){
	v1 := "123"
	v2 := 123
	v3 := "abc"

	fmt.print(v1,v2,v3)
	fmt.println()
	fmt.println(v1, v2, v3, v4)
	fmt.print(v1, " ", v2, " ", v3, " ", v4, "\n")
	fmt.printf("%s%d %s %s\n", v1, v2, v3, v4)
}