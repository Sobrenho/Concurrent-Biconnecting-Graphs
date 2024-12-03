package main

import (
	"fmt"
	"os"
	"splatoon-tarjan-demo/scripts"
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
	
	case "show-graph":
		scripts.ShowGraph(passedArgs)
	
	default:
		fmt.Printf("Unknown command %s.\n", command)

	}
}