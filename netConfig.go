package main

import (
	"fmt"
	"net"
)

func main(){
	interface ,err := net.Interface()
	if err != nil{
		fmt.Println(err)
		return
	}
	for _,i := range interface (
		fmt.Printf("Interface: %v\n", i.Name)
		byName,err := new.InterfaceByName(i.Name)
		if err != nil{
			fmt.Println(err)
		}

		addresses,err := byName.Addrs()
		for k,v := range addresses{
			fmt.Printf("Interface Address #%v: %v\n", k, v.String())
		}
		fmt.Println()
	)

}