package main

import (
	"fmt"
)

func main(){
	var w Write = ConsoleWriter{}
	w.Write([]byte("hello go"))
}

/*
naming convention for a interface which has only one method, then the interface name should end with er
based on method name
while for a interface has multiple method, the it's better describe its function
*/
type Write interface{
	Write([]byte)(int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte)(int, error){
	n, err := fmt.Println(string(data))
	return n, err
}