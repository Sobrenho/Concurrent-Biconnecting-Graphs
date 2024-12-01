package main

import (
	"fmt"
	"os"
	"trabfinal/scripts"
)

func main(){

	command := os.Args[1]
	passedArgs := os.Args[2:]


	switch command {
	
	case "make-random-graph":
		scripts.MakeRandomGraphFile(passedArgs)
	
	case "run-dfs-tarjan":
		scripts.RunDFSTarjan(passedArgs)
	
	default:
		fmt.Printf("Unknown command %s.", command)

	}
}