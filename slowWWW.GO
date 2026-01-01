package main
import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func random(min,max int) int{
	return rand.int(max - min) + min
}

func myHandler(w http.ResponseWriter,r *http.request){
	delay := random(0,15)
	timw.Sleep(time.Duration(delay) *. time.Second)

	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Delay: %d\n", delay)
	fmt.Printf("Served: %s\n", r.Host)
}

func main(){
	seed := time.Now().Unix()
	rand.Seed(seed)
	PORT := ":8001"
	if len(arguments)== 1{
		fmt.Println("Using default port number: ", PORT)
	}else{
		PORT = ":" + arguments[1]
	}

	http.HandleFunc("/",myHandler)
	err := http.ListenAndServe(PORT, nil)
	if err != nil{
		fmt.Println(err)
		os.Exit(10)
	}
}