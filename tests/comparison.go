package tests

import (
	"encoding/csv"
	"fmt"
	"os"
	"splatoon-tarjan-demo/graphs"
	"strconv"
	"time"
)




func CompareImplementations(TOTAL_GRAPHS int){
	sequencial := make([]int64, TOTAL_GRAPHS)
	concorrente_1T := make([]int64, TOTAL_GRAPHS)
	concorrente_2T := make([]int64, TOTAL_GRAPHS)
	concorrente_4T := make([]int64, TOTAL_GRAPHS)
	concorrente_8T := make([]int64, TOTAL_GRAPHS)

	for i := range TOTAL_GRAPHS{
		//Gerar Grafo
		G := graphs.NewRandomGraph(1000, 2000)

		//Versão Sequencial
		t_seq := time.Now()
		G.DFSTarjan()
		sequencial[i] = time.Since(t_seq).Microseconds()

		// 1 Thread
		t_1t := time.Now()
		G.SplatoonTarjan(1)
		concorrente_1T[i] = time.Since(t_1t).Microseconds()

		// 2 Thread
		t_2t := time.Now()
		G.SplatoonTarjan(2)
		concorrente_2T[i] = time.Since(t_2t).Microseconds()

		// 4 Thread
		t_4t := time.Now()
		G.SplatoonTarjan(4)
		concorrente_4T[i] = time.Since(t_4t).Milliseconds()

		// 8 Thread
		t_8t := time.Now()
		G.SplatoonTarjan(8)
		concorrente_8T[i] = time.Since(t_8t).Milliseconds()
	}

	file, err := os.Create("timeTable.csv")
	if err != nil{
		fmt.Println("Erro ao criar csv.")
		return
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	w.Write([]string{"ID_Grafo", "Tempo Sequencial","Tempo 1 Thread", "Tempo 2 Threads", "Tempo 4 Threads", "Tempo 8 Threads"})

	for i := range TOTAL_GRAPHS{
		w.Write([]string{strconv.Itoa(i), 
						 strconv.FormatInt(sequencial[i], 10),
						 strconv.FormatInt(concorrente_1T[i],10),
						 strconv.FormatInt(concorrente_2T[i], 10),
						 strconv.FormatInt(concorrente_4T[i], 10), 
						 strconv.FormatInt(concorrente_8T[i], 10),
						},
				)
	}

}