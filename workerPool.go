package main
import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type client struct{
	id int
	integer int
}

type Data struct {
	job client
	square int
}

var {
	size = 10
	clients = make(chan chan,size)
	data = make(chan Data,size)
}

func worker(w *sync.WaitGroup){
	for c := range clients{
		square := c.integer * c.integer
		output := Data{c, square}
		data <- output
		time.Sleep(time.Second)
	}
	w.Done()
}
func makeWp(n int){
	var w sync.WaitGroup
	for i := 0;i < n;i++{
		w.Add(1)
		go.worker(&w)
	}
	w.wait()
	close(data)
}

func create(n int){
	for i := 0; i< n; i++{
		c := Client{i,i}
		clients <- c
	}
	close(clients)
}



func main(){
	fmt.Println("Capacity of clients:", cap(clients))
	fmt.Println("Capacity of data:", cap(data))

	if len(os.Args) != 3 {
		fmt.Println("Need #jobs and #workers!")
		os.Exit(1)
	}

	nJobs ,err := strconv.Atoi(os.Args[2])
	if err != nil{
		fmt.Println(err)
		return		
	}
	nWorkers,err := strconv.Atoi(os.Args[])
	if err != nil{
		fmt.Println(err)
		return
	}
}