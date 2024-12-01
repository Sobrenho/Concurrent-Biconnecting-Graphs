package graphs

import (
	"sync"
	"trabfinal/unionfind"
)

func (graph *GraphX) Splatoon(threadsCount int) []int {

	nextVertex := 0
	var nextVertexLock sync.Mutex

	isVisited := make([]bool, graph.VerticesCount())

	unionFind := unionfind.NewUnionFind(graph.VerticesCount())

	returnedChannel := make(chan bool, threadsCount)

	for i := 0; i < threadsCount; i++ {

		go func() {

			for {

				nextVertexLock.Lock()
				if nextVertex == graph.VerticesCount() {
					nextVertexLock.Unlock()
					break
				}
				vertex := nextVertex
				nextVertex++
				nextVertexLock.Unlock()

				isVisited[vertex] = true
	
				for _, neighbor := range graph.Adjacents(vertex) {
					unionFind.Join(vertex, neighbor)
				}
			}

			returnedChannel <- true
	
		}()

	}

	for i := 0; i < threadsCount; i++ {
		<- returnedChannel
	}
	return unionFind.Representatives()
}
