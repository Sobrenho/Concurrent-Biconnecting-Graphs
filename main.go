package main

import (
	"fmt"
	"os"
	"splatoon-tarjan-demo/scripts"
	"splatoon-tarjan-demo/tests"
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
	
	case "measure-execution-time": // Como rodar .\executavel measure-execution-time <Número de iterações por grafo>
		tests.MeasureRunningTime(passedArgs)
	default:
		fmt.Printf("Unknown command %s.\n", command)

	}
}