package graphs

type Edge struct {
	V1 *Vertex
	V2 *Vertex
}


type Vertex struct {
	Id		int
	Color  	string
	Parent 	*Vertex
	Desc    int //Tempo de Descoberta
	Fim     int //Tempo de Finalização
	Ret    	*int //Vértice de menor tempo de descoberta dentre os filhos
}

type Graph struct {
	N 				int
	M 				int
	Vertices		[]Vertex
	AdjacencyList 	map[*Vertex][]*Vertex
}


func (G *Graph) AmostrarVert() *Vertex{
	for v := range G.AdjacencyList{
		return v
	}
	
	return nil
}

// Gera seguinte grafo: https://i.imgur.com/8s24EVp.png
func GenGraph() (*Graph) {
	G := new(Graph)
	G.AdjacencyList = make(map[*Vertex][]*Vertex, 17)

	G.N = 18
	G.M = 19
	
	G.Vertices = make([]Vertex, 0)

	//Criação dos objetos vértice
	for i := range G.N{
		G.Vertices = append(G.Vertices, Vertex{i, "WHITE", nil,0,0, new(int)})
	}
	

	//Criação das arestas
	G.AdjacencyList[&G.Vertices[0]] 	= []*Vertex{&G.Vertices[1],  &G.Vertices[2]} 
	G.AdjacencyList[&G.Vertices[1]] 	= []*Vertex{&G.Vertices[0],  &G.Vertices[3]}
	G.AdjacencyList[&G.Vertices[2]] 	= []*Vertex{&G.Vertices[0],  &G.Vertices[3], &G.Vertices[4]}
	G.AdjacencyList[&G.Vertices[3]] 	= []*Vertex{&G.Vertices[1],  &G.Vertices[2], &G.Vertices[4]}
	G.AdjacencyList[&G.Vertices[4]] 	= []*Vertex{&G.Vertices[2],  &G.Vertices[3], &G.Vertices[5]}
	G.AdjacencyList[&G.Vertices[5]] 	= []*Vertex{&G.Vertices[4]}
	G.AdjacencyList[&G.Vertices[6]] 	= []*Vertex{&G.Vertices[7],  &G.Vertices[8]}
	G.AdjacencyList[&G.Vertices[7]] 	= []*Vertex{&G.Vertices[6],  &G.Vertices[8]}
	G.AdjacencyList[&G.Vertices[8]] 	= []*Vertex{&G.Vertices[6],  &G.Vertices[7]}
	G.AdjacencyList[&G.Vertices[9]] 	= []*Vertex{&G.Vertices[10], &G.Vertices[11]}
	G.AdjacencyList[&G.Vertices[10]] 	= []*Vertex{&G.Vertices[9],  &G.Vertices[11], &G.Vertices[12]}
	G.AdjacencyList[&G.Vertices[11]] 	= []*Vertex{&G.Vertices[9],  &G.Vertices[10], &G.Vertices[12]}
	G.AdjacencyList[&G.Vertices[12]] 	= []*Vertex{&G.Vertices[10], &G.Vertices[11]}
	G.AdjacencyList[&G.Vertices[13]] 	= []*Vertex{&G.Vertices[14], &G.Vertices[16]}
	G.AdjacencyList[&G.Vertices[14]] 	= []*Vertex{&G.Vertices[13], &G.Vertices[15]}
	G.AdjacencyList[&G.Vertices[15]] 	= []*Vertex{&G.Vertices[14], &G.Vertices[16]}
	G.AdjacencyList[&G.Vertices[16]] 	= []*Vertex{&G.Vertices[13], &G.Vertices[15]}
	G.AdjacencyList[&G.Vertices[17]] 	= nil

	return G

}



func TesteBiconnect() (*Graph){
	G := new(Graph)
	G.AdjacencyList = make(map[*Vertex][]*Vertex, 6)

	G.N = 6
	G.M = 7
	
	var Vertices []Vertex

	//Criação dos objetos vértice
	for i := range G.N{
		Vertices = append(Vertices, Vertex{i, "WHITE", nil,0,0, new(int)})
	}
	

	//Criação das arestas
	G.AdjacencyList[&Vertices[0]] 	= []*Vertex{&Vertices[1], &Vertices[2]} 
	G.AdjacencyList[&Vertices[1]] 	= []*Vertex{&Vertices[0], &Vertices[3]}
	G.AdjacencyList[&Vertices[2]] 	= []*Vertex{&Vertices[0], &Vertices[3], &Vertices[4]}
	G.AdjacencyList[&Vertices[3]] 	= []*Vertex{&Vertices[1], &Vertices[2], &Vertices[4]}
	G.AdjacencyList[&Vertices[4]] 	= []*Vertex{&Vertices[2], &Vertices[3], &Vertices[5]}
	G.AdjacencyList[&Vertices[5]] 	= []*Vertex{&Vertices[4]}

	return G
}

