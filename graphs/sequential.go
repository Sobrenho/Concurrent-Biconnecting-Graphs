package graphs

import "fmt"

func (G *Graph) DFS_Sequencial() {

	time := 0
	for u := range G.AdjancencyList {
		u.Color = "WHITE"
		u.Parent = nil
	}

	for u := range G.AdjancencyList {
		if u.Color == "WHITE" {
			G.dfs_Visit(u, &time)
		}
	}

}

func (G *Graph) dfs_Visit(u *Vertex, time *int) { //DFS_Visit
	*time += 1
	u.Desc = *time
	u.Color = "GRAY"

	for _, v := range G.AdjancencyList[u] {
		if v.Color == "WHITE" {
			v.Parent = u
			G.dfs_Visit(v, time)
		}
	}
		u.Color = "BLACK"
		*time += 1
		u.Fim = *time
}


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