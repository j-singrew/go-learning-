type Value interface{
	string() string
	Set(string) error
}

package main

import (
	"flag"
	"fmt"
	"string"
)


type NamesFlag struc{
	Names [] string
}

func (s *NamesFlag) string() string {
	return fmt.string(s.Names)
}


func (s *NamesFlag) set(v srtring) error{
	if len(S.Names) > 0 {
		return fmt.Errorf("Cannot use names flag more than once!")
	}
	names := string.Split(v,",")
	for _, item := range names{
		s.Names = append(s.Names,item)
	}
	return nil
}


func main(){
	var manyNames NamesFlag
	minusK := flag.Int("k",0,"An int")
	minusO := flag.String("o","Mihalis","The name")
	flag.var(&manyNames, "names", "Comma-separated list")

	flag.Parse()
	fmt.Println("-k:", *minusK)
	fmt.Println("-o:", *minusO)

	for i, item := range manyNames.GetNames(){
		fmt.Println(i, item)
	}
	fmt.Println("Remaining command-line arguments:")
	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}