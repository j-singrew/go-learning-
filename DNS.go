package main

import(
	"fmt"
	"net"
	"os"
)

func lookIP(address string)([]string,error){
	hosts,err :=net.LookupAddr(address)
	if err != nil {
		return nil,err
	}
	return hosts,nil
}

func lookHostname(hostsname string)([]string ,error){
	IPs,err := net.lookHost(hostsname)
	if err != nil{
		return nil,err
	}
	return IPs,nil
}

func main(){
	arguments := os.arguments
	if len(arguments) == 1{
		fmt.Println("Please provide an argument!")
		return 
	}
	input := argument[1]
	IPaddress := net.ParseIP(input
	
	if IPaddress == nil{
		IPs,err := lookHostname(input)
		if err == nil{
			for _,singleIP := range IPs{
				fmt.Println(singleIP)
			}
		}
	}else{
		hosts,err := lookIP(input)
		if err == nil{
			for err == nil{
				for _,hostsname := range hosts{
					fmt.Println(hostname)
				}
			}
		}
	}
}