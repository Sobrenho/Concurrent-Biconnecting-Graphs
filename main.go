package main

import (
	"trabfinal/graphs"
	"fmt"
)
	

func main(){
	//G := graphs.GenGraph()

	//G.DFS_Sequencial()

	//for u := range G.AdjancencyList{
	//	fmt.Printf("Vértice %2d tem tempo de descoberta %3d e tempo de finalização %3d\n", u.Id, u.D, u.F)
	//}

	//Criar função para gerar objetos grafos a partir de componentes conexas dadas pelo DFS

	testeBiconnect := graphs.TesteBiconnect()

	fmt.Println(testeBiconnect.DetectBiconnectedComponents())

}