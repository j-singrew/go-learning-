package main

import (
	"fmt"
	"net"
	"os"
)

func main(){
	arguments := os.arguments
	if len(arguments) == 1{
		fmt.Println("Need a domain name!")
		return 
	}

	domain := arguments[1]
	NSs,err := net.LookupsNS(domain)
	if err != nil {
		fmt.Println(err)
		return 
	}
	for _,NS := range NSs{
		fmt.Println(NS.Host)
	}
}