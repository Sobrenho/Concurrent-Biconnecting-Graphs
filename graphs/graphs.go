package graphs


type Edge struct {
	V1 *Vertex
	V2 *Vertex
}

type Vertex struct {
	Id		int
	Color  	string
	Parent 	*Vertex
	Desc      	int //Tempo de Descoberta
	Fim     	int //Tempo de Finalização
	Ret    	*int //Vértice de menor tempo de descoberta dentre os filhos
}

type Graph struct {
	N 				int
	M 				int
	AdjancencyList 	map[*Vertex][]*Vertex
}


func (G *Graph) AmostrarVert() *Vertex{
	for v := range G.AdjancencyList{
		return v
	}
	
	return &Vertex{}
}

// Gera seguinte grafo: https://i.imgur.com/8s24EVp.png
func GenGraph() (*Graph) {
	G := new(Graph)
	G.AdjancencyList = make(map[*Vertex][]*Vertex, 17)

	G.N = 18
	G.M = 19
	
	var Vertices []*Vertex

	//Criação dos objetos vértice
	for i := range G.N{
		Vertices = append(Vertices, &Vertex{i, "WHITE", nil, 0,0, new(int)})
	}
	

	//Criação das arestas
	G.AdjancencyList[Vertices[0]] 	= []*Vertex{Vertices[1], Vertices[2]} 
	G.AdjancencyList[Vertices[1]] 	= []*Vertex{Vertices[0], Vertices[3]}
	G.AdjancencyList[Vertices[2]] 	= []*Vertex{Vertices[0], Vertices[3], Vertices[4]}
	G.AdjancencyList[Vertices[3]] 	= []*Vertex{Vertices[1], Vertices[2], Vertices[4]}
	G.AdjancencyList[Vertices[4]] 	= []*Vertex{Vertices[2], Vertices[3], Vertices[5]}
	G.AdjancencyList[Vertices[5]] 	= []*Vertex{Vertices[4]}
	G.AdjancencyList[Vertices[6]] 	= []*Vertex{Vertices[7], Vertices[8]}
	G.AdjancencyList[Vertices[7]] 	= []*Vertex{Vertices[6], Vertices[8]}
	G.AdjancencyList[Vertices[8]] 	= []*Vertex{Vertices[6], Vertices[7]}
	G.AdjancencyList[Vertices[9]] 	= []*Vertex{Vertices[10], Vertices[11]}
	G.AdjancencyList[Vertices[10]] 	= []*Vertex{Vertices[9], Vertices[11], Vertices[12]}
	G.AdjancencyList[Vertices[11]] 	= []*Vertex{Vertices[9], Vertices[10], Vertices[12]}
	G.AdjancencyList[Vertices[12]] 	= []*Vertex{Vertices[10], Vertices[11]}
	G.AdjancencyList[Vertices[13]] 	= []*Vertex{Vertices[14], Vertices[16]}
	G.AdjancencyList[Vertices[14]] 	= []*Vertex{Vertices[13], Vertices[15]}
	G.AdjancencyList[Vertices[15]] 	= []*Vertex{Vertices[14], Vertices[16]}
	G.AdjancencyList[Vertices[16]] 	= []*Vertex{Vertices[13], Vertices[15]}

	return G

}



func TesteBiconnect() (*Graph){
	G := new(Graph)
	G.AdjancencyList = make(map[*Vertex][]*Vertex, 6)

	G.N = 6
	G.M = 7
	
	var Vertices []*Vertex

	//Criação dos objetos vértice
	for i := range G.N{
		Vertices = append(Vertices, &Vertex{i, "WHITE", nil, 0,0, new(int)})
	}
	

	//Criação das arestas
	G.AdjancencyList[Vertices[0]] 	= []*Vertex{Vertices[1], Vertices[2]} 
	G.AdjancencyList[Vertices[1]] 	= []*Vertex{Vertices[0], Vertices[3]}
	G.AdjancencyList[Vertices[2]] 	= []*Vertex{Vertices[0], Vertices[3], Vertices[4]}
	G.AdjancencyList[Vertices[3]] 	= []*Vertex{Vertices[1], Vertices[2], Vertices[4]}
	G.AdjancencyList[Vertices[4]] 	= []*Vertex{Vertices[2], Vertices[3] , Vertices[5]}
	G.AdjancencyList[Vertices[5]] 	= []*Vertex{Vertices[4]}

	return G
}

