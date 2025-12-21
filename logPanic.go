package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main(){
	sysLog,err := syslog.New(sysLog.LOG_ALERT|sysLog.LOG_MAIL,"some program!")

	if err != nil {
		log.Fatal(err)
	}else{
		log.SetOutput(sysLog)
	}
	log.Panic(sysLog)
	fmt.Println("will you see this ?")
}