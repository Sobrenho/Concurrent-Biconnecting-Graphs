package main

import (
	"fmt"
	"os"
	"trabfinal/scripts"
)

func main(){

	switch os.Args[1] {
	
	case "make-random-graph":
		scripts.MakeRandomGraphFile(os.Args[2:])
	
	default:
		fmt.Printf("Unknown command %s.", os.Args[1])

	}
}