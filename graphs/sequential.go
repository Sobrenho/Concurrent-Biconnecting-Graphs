package graphs

import (
	"slices"
)



// ----- DFS ----------------

func (G *Graph) DFS_Sequencial() []*Graph{
	connectedComponents := make([]*Graph, 0, G.N)



	for v := range G.AdjacencyList{
		v.Color = "WHITE"
		v.Parent = nil
	}

	//Construir Lista de Adjacencia


	time := 0
	
	for u := range G.AdjacencyList {
		
		if (*u).Color == "WHITE" {

			treeVertices := G.dfs_Visit(u, &time)
			componentAdjacencyList := make(map[*Vertex][]*Vertex)

			edges := 0 

			for _, v := range treeVertices{

				for _, w := range G.AdjacencyList[v]{
					if slices.Contains(treeVertices, w){
						componentAdjacencyList[v] = append(componentAdjacencyList[v], w)
						edges++
					}
				}
			}

			//Transformar []*Vertex em []Vertex
			new_graph_vertices := make([]Vertex, 0, len(treeVertices))
			for _, ptr := range treeVertices{
				new_graph_vertices = append(new_graph_vertices, *ptr)
			}

			connectedComponents = append(connectedComponents, &Graph{len(treeVertices), edges/2 , new_graph_vertices, componentAdjacencyList})

		}
	}
	return connectedComponents
}


//Função Recursiva DFS 
//Retorna uma lista com todos os vértices visitados abaixo do vértice u
func (G *Graph) dfs_Visit(u *Vertex, time *int) []*Vertex{ 
	*time += 1
	u.Desc = *time
	u.Color = "GRAY"

	visitedVertices := make([]*Vertex, 0, G.N)
	
	for i := range G.AdjacencyList[u] {
		v := G.AdjacencyList[u][i]
		if v.Color == "WHITE" {
			v.Parent = u

			visitedVertices = append(visitedVertices, G.dfs_Visit(v, time)...)

		}
	}
		u.Color = "BLACK"
		*time += 1
		u.Fim = *time

		return append(visitedVertices, u)
}

// -----------------------------------------




// ------ Tarjan Componente Biconexa -------------

func (G * Graph) DetectBiconnectedComponents() bool{
	tempo := 0
	edge_stack := make(stack[Edge], G.M)

	for u := range G.AdjacencyList {
		u.Color = "WHITE"
		u.Parent = nil
	}
	
	v := G.AmostrarVert()
	if v == nil{ //Grafo possui um único vértice
		return false
	}

	arestas_no_bloco := G.detectBiconnectedComponentsVisit(v, &edge_stack, &tempo)


	return arestas_no_bloco == G.M
}


func (G * Graph) detectBiconnectedComponentsVisit(u *Vertex, pilha *stack[Edge], tempo *int) int{
	var aresta Edge
	
	u.Color = "GRAY"
	*tempo +=  1
	u.Desc = *tempo
	*u.Ret = u.Desc
	
	block_size := 0


	for _, v := range G.AdjacencyList[u] {
		if v.Color == "WHITE"{
			v.Parent = u
			*pilha = pilha.Push(Edge{u,v})

			block_size = G.detectBiconnectedComponentsVisit(v, pilha, tempo)

			if *v.Ret >= u.Desc { //Achou articulação 
				edge_count := 0
				for {
					aresta, *pilha = pilha.Pop()
					edge_count++

					if aresta.V1 == u && aresta.V2 == v {
						return edge_count
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
	return block_size
}