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
		scripts.MakeRandomGraph(passedArgs)
	
	case "run-dfs-tarjan":
		scripts.RunDFSTarjan(passedArgs)
	
	case "run-splatoon-tarjan":
		scripts.RunSplatoonTarjan(passedArgs)
    
    case "validate-splatoon-tarjan":
        scripts.ValidateSplatoonTarjan(passedArgs)
	
	default:
		fmt.Printf("Unknown command %s.\n", command)

	}
}