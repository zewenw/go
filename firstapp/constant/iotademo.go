package main

import (
	"fmt"
)

const (
	_ = iota // "_" means ignore first value by assigning to blank identitier

	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main(){
	fileSize := 4000000000.
	fmt.Printf("%.3f GB", fileSize / GB)
}