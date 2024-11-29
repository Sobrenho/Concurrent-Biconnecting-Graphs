package graphs

import (
	"trabfinal/unionfind"
)

func (graph *GraphX) Splatoon(threadsCount int) []int {

	verticesChannel := make(chan int, graph.VerticesCount())
	for u := 0; u < graph.VerticesCount(); u++ {
		verticesChannel <- u
	}
	
	canFinish := make(chan bool, 1)

	verticesConsumed := make(chan int, 1)
	verticesConsumed <- 0

	isVisited := make([]bool, graph.VerticesCount())
	unionFind := unionfind.NewUnionFind(graph.VerticesCount())

	threadsFinished := make(chan int, 1)
	threadsFinished <- 0

	canReturn := make(chan bool)

	for i := 0; i < threadsCount; i++ {

		go func() {

			mustContinue := true
			
			for mustContinue {
				select {
	
				case vertex := <- verticesChannel:
	
					isVisited[vertex] = true
	
					for _, neighbor := range graph.Adjacents(vertex) {
						unionFind.Join(vertex, neighbor)
					}
	
					vConsumed := <- verticesConsumed
					vConsumed++
					verticesConsumed <- vConsumed
	
					if vConsumed == graph.VerticesCount() {
						mustContinue = false
					}
	
				case <- canFinish:
					mustContinue = false
				}
	
			}

			tFinished := <- threadsFinished
			tFinished++
			threadsFinished <- tFinished

			canFinish <- true

			if tFinished == threadsCount {
				canReturn <- true
			}
	
		}()

	}

	<- canReturn
	return unionFind.Representatives()
}

func (graph *Graph) SplatoonComponentSearch(threadsCount int) []*Vertex {

	verticesChannel := make(chan *Vertex, graph.N)
	for i := 0; i < graph.N; i++ {
		verticesChannel <- &graph.Vertices[i]
	}
	
	canFinish := make(chan bool)

	verticesConsumed := make(chan int, 1)
	verticesConsumed <- 0

	isVisited := make(map[*Vertex]bool, graph.N)

	unionFind := unionfind.NewUnionFind(graph.N)

	for i := 0; i < threadsCount; i++ {

		go func(verticesChannel chan *Vertex, canFinish chan bool, verticesConsumed chan int) {

			for {
				select {
	
				case vertex := <- verticesChannel:
	
					isVisited[vertex] = true
	
					for _, neighbor := range graph.AdjacencyList[vertex] {
						unionFind.Join(vertex.Id, neighbor.Id)
					}
	
					vConsumed := <- verticesConsumed
					
					vConsumed++
	
					verticesConsumed <- vConsumed
	
					if vConsumed == graph.N {
						canFinish <- true
						return
					}
	
				case <- canFinish:
	
					canFinish <- true
					return
				
				}
	
			}
	
		}(verticesChannel, canFinish, verticesConsumed)

	}

	<- canFinish

	representatives := unionFind.Representatives()
	representativeVertices := make([]*Vertex, len(representatives))

	for i, representative := range representatives {
		representativeVertices[i] = &graph.Vertices[representative]
	}

	return representativeVertices
}