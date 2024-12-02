package tests

import (
	"encoding/csv"
	"fmt"
	"os"
	"log"
	"splatoon-tarjan-demo/scripts"
	"strconv"
	"strings"
)



func createCSV(timestamps [][]int64){

	f, err := os.Create("timeTable.csv")
	if err != nil{
		fmt.Fprintln(os.Stderr, "Erro ao criar csv")
		return
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()
	w.Write([]string{"grafo", "Sequencial", "1 Thread", "2 Threads", "4 Threads","8 Threads", "16 Threads"})

	for row := range timestamps{
		w.Write([]string{strconv.Itoa(row), // Identificador do grafo
				strconv.FormatInt(timestamps[row][0], 10),
				strconv.FormatInt(timestamps[row][1], 10),
				strconv.FormatInt(timestamps[row][2], 10),
				strconv.FormatInt(timestamps[row][3], 10),
				strconv.FormatInt(timestamps[row][4], 10),
				strconv.FormatInt(timestamps[row][5], 10),
				})
	}
}



func MeasureRunningTime(args []string){
	//Constantes para logging
	colorReset := "\033[0m"
	colorGreen := "\033[32m"

	iterationsPerGraphSizes, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
		return 
	}

	
	graph_sizes := []string{ "1000 500000",
    						 "5000 12500000",
    						 "10000 50000000",
							}

	timeTable := make([][]int64, len(graph_sizes))
	for row := range timeTable{
		timeTable[row] = make([]int64, 6)
	}
	

	for i, g_size := range graph_sizes{

		fmt.Printf("Gerando grafo: %s...", g_size)
		scripts.MakeRandomGraph(append(strings.Split(g_size, " "), "g"))
		fmt.Println(colorGreen, "done", colorReset)

		fmt.Print("Sequencial...")
		timeTable[i][0] = mean(iterationsPerGraphSizes, []string{"g", "out"}, scripts.RunDFSTarjan)
		fmt.Println(colorGreen, "done", colorReset)


		for j, numThreads := range []int{1,2,4,8,16}{
			fmt.Printf("Concorrente - %d thread....", numThreads)
			timeTable[i][j+1] = mean(iterationsPerGraphSizes, []string{"g", strconv.Itoa(numThreads), "out"}, scripts.RunSplatoonTarjan)
			fmt.Println(colorGreen, "done", colorReset)
		}
	}

	createCSV(timeTable)
}

func mean(iter int, args []string , exec func([]string) int64) int64{
	var sum int64
	for i := 0; i < iter ; i++{
		sum += exec(args)
	}

	return sum/int64(iter)
}



