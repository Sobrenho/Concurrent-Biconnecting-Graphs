package graphs


// ---------- DFS ----------------------------------------

func (G *Graph) DFS_Sequencial() []*Vertex{
	sourcesForComponents := make([]*Vertex, 0, G.N)


	for v := range G.AdjacencyList{
		v.Color = "WHITE"
		v.Parent = nil
	}


	time := 0
	
	for u := range G.AdjacencyList {
		
		if (*u).Color == "WHITE" {
			G.dfs_Visit(u, &time)
			sourcesForComponents = append(sourcesForComponents, u)
		}
	}

	return sourcesForComponents
}


/*
Função Recursiva DFS 

- Retorna uma lista com todos os vértices visitados abaixo do vértice u

- Descomentar linhas com sifrão ($) para retornar todos os vértices da componente

*/
func (G *Graph) dfs_Visit(u *Vertex, time *int) {//$ []*Vertex{
	*time += 1
	u.Desc = *time
	u.Color = "GRAY"

	//$visitedVertices := make([]*Vertex, 0, G.N)
	
	for i := range G.AdjacencyList[u] {
		v := G.AdjacencyList[u][i]
		if v.Color == "WHITE" {
			v.Parent = u
			G.dfs_Visit(v, time)
			//$ visitedVertices = append(visitedVertices, G.dfs_Visit(v, time)...)
		}
	}
		u.Color = "BLACK"
		*time += 1
		u.Fim = *time

		//$ return append(visitedVertices, u)
}

// -----------------------------------------




// ------ Tarjan Componente Biconexa -------------

func (G * Graph) CheckIfBiconnectedComponentWithSourceVertex(source *Vertex) bool{
	tempo := 0
	edge_stack := make(stack[Edge], G.M)

	for u := range G.AdjacencyList {
		u.Color = "WHITE"
		u.Parent = nil
	}
	
	numberOfBlocks := G.checkIfBiconnectedComponentWithSourceVertexVisit(source, &edge_stack, &tempo)

	return numberOfBlocks == 1
}


func (G * Graph) checkIfBiconnectedComponentWithSourceVertexVisit(u *Vertex, pilha *stack[Edge], tempo *int) int{
	var aresta Edge
	
	u.Color = "GRAY"
	*tempo +=  1
	u.Desc = *tempo
	*u.Ret = u.Desc
	
	num_blocks := 0


	for _, v := range G.AdjacencyList[u] {
		if v.Color == "WHITE"{
			v.Parent = u
			*pilha = pilha.Push(Edge{u,v})

			num_blocks = G.checkIfBiconnectedComponentWithSourceVertexVisit(v, pilha, tempo)

			if *v.Ret >= u.Desc { //Achou articulação 
				for {
					aresta, *pilha = pilha.Pop()
					if aresta.V1 == u && aresta.V2 == v {
						num_blocks++
						break
					}
				}
			}else{
				*u.Ret = min(*u.Ret, *v.Ret)
			}

		}else{
			if v != u.Parent{
				if v.Desc < u.Desc {
					*pilha = pilha.Push(Edge{u,v})
				}

				*u.Ret = min(*u.Ret, v.Desc)
			}
		}

	}
	return num_blocks
}



func (G *Graph)BiconnectedComponentsInGraphCheck () ([]*Vertex, []bool){
	componentesConexas := G.DFS_Sequencial()
	isBiconnected := make([]bool, len(componentesConexas))

	for i, v := range componentesConexas{
		isBiconnected[i] = G.CheckIfBiconnectedComponentWithSourceVertex(v)
	}

	return componentesConexas, isBiconnected
}
