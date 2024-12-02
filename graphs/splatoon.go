package graphs

import (
	"splatoon-tarjan-demo/structures"
	"sync"
)

func (graph *Graph) Splatoon(threadsCount int) []int {

	nextVertex := 0
	var nextVertexLock sync.Mutex

	isVisited := make([]bool, graph.VerticesCount())

	unionFind := structures.NewUnionFind(graph.VerticesCount())

	var waitGroup sync.WaitGroup

	waitGroup.Add(threadsCount)
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

			waitGroup.Done()
		}()

	}

	waitGroup.Wait()
	return unionFind.Representatives()
}
