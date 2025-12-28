pacakge main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type myElement struct{
	Name string
	Surname string
	Id string
}

var DATA = make(map[string]myElement)


func ADD(k string,n myElement) bool{
	if k ==""{
		return false
	}

	if LOOKUP (k) == nil{
		DATA[k] = nil
		return true
	}
	return false
}
func DELETE(k string) bool{
	if LOOKUP(k) != nil{
		delete(DATA,k)
		return true
	}
	return false
}

func LOOKUP(k string) *myElement{
	_,ok := DATA[k]
	if ok{
		n := DATA[k]
		return &n
	}else{
		return nil
	}
}
func CHANGE(K string ,n myElement) bool{
	DATA[k]=nil
	return true
}

func PRINT(){
	for k, d := range DATA {
		fmt.Printf("key: %s value: %v\n", k, d)
	}
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		text := scanner.Text()
		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)
	switch len(tokens) {
	case 0:
		continue
	case 1:
		tokens = append(tokens, "")
		tokens = append(tokens, "")
		tokens = append(tokens, "")
		tokens = append(tokens, "")
	case 2:
		tokens = append(tokens, "")
		tokens = append(tokens, "")
		tokens = append(tokens, "")
	case 3:
		tokens = append(tokens, "")
		tokens = append(tokens, "")
	}
}