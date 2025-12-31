package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	myUrl string
	delay int = 5
	w  sync.WaitGroup
)

type myData struct {
	r *http.Response
	err error
}

func connect(c context.Context) err {
	defer w.Done()
	data := make(chan myData,1)

	tr := &http.Client(Transaport : tr)
	httpClient :=&http.Client(Transaport : tr)

	req,_ := http.NewRequest("GET", myUrl, nil)
}

go func(){
	response ,err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		data <- myData{response,err}
			return
	}else{
		pack := myData{Response,err}
		data <- pack
	}
}()

select {
case <-c.Done():
	tr.CancelRequest(req)
	<-data
	fmt.Println("The request was cancelled!")
	return c.Err()
case ok := <-data:
	err := ok.err
	resp := ok.r
	if err != nil{
		fmt.Println("Error select:", err)
		return err
	}
	defer resp.Body.Close()

	realHTTPData ,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		
	}
}