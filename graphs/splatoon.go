package graphs

import (
	"sync"
	"trabfinal/unionfind"
)

func (graph *GraphX) Splatoon(threadsCount int) []int {

	verticesChannel := make(chan int, graph.VerticesCount())
	for u := 0; u < graph.VerticesCount(); u++ {
		verticesChannel <- u
	}
	
	canFinish := make(chan bool, 1)

	verticesConsumed := 0
	var verticesConsumedLock sync.Mutex

	isVisited := make([]bool, graph.VerticesCount())

	unionFind := unionfind.NewUnionFind(graph.VerticesCount())

	threadsFinished := 0
	var threadsFinishedLock sync.Mutex

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
	
					verticesConsumedLock.Lock()
					verticesConsumed++
					if verticesConsumed == graph.VerticesCount() {
						mustContinue = false
					}
					verticesConsumedLock.Unlock()
	
				case <- canFinish:
					mustContinue = false
				}
	
			}

			threadsFinishedLock.Lock()
			threadsFinished++
			if threadsFinished == threadsCount {
				canReturn <- true
			}
			canFinish <- true
			threadsFinishedLock.Unlock()
	
		}()

	}

	<- canReturn
	return unionFind.Representatives()
}
