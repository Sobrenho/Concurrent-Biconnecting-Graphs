package graphs

import (
	"sync"
)


func (graph *Graph) ShiloachVishkin(threadsCount int) []int{


	parent := make([]int, graph.VerticesCount())

	vertChanHooking := make(chan int, graph.VerticesCount())
	vertChanCompressing := make(chan int, graph.VerticesCount())

	endHooking := make(chan struct{}, 1) //Faz o papel de mutex
	endCompressing := make(chan struct{}, 1) //Faz o papel de mutex

	for i := 0; i < graph.VerticesCount(); i++ {
		parent[i] = i
		vertChanHooking <- i
	}
	
	update := make(chan bool, 1)

	var usedVertices int
	var usedVerticesLock sync.Mutex

	var wgThreads sync.WaitGroup
	wgThreads.Add(threadsCount)

	for <-update{

		for i := 0; i< threadsCount; i++{
			//Hooking
			go func(){
				for{
					select{
					case v:= <-vertChanHooking:

						usedVerticesLock.Lock()
						usedVertices++
						usedVerticesLock.Unlock()

						vertChanCompressing <- v

						cur_parent := parent[v]
						for _, neighbor := range graph.Adjacents(v){
							if parent[neighbor] < cur_parent{

								parent[v] = parent[neighbor]

								<-update
								update<-true

							}

						}

						usedVerticesLock.Lock()
						if usedVertices == graph.VerticesCount(){
							endHooking <- struct{}{}
						}
						usedVerticesLock.Unlock()

					case <- endHooking:
						endHooking <- struct{}{}
						wgThreads.Done()

					}
				}
			}()
		}
			wgThreads.Wait()
			
			usedVertices = 0


		for i:=0; i< threadsCount; i++{
			//Compressing
			go func() {
				select{

				case v := <- vertChanCompressing:

					usedVerticesLock.Lock()
					usedVertices++
					usedVerticesLock.Unlock()

					vertChanHooking <- v

					for parent[parent[v]] != parent[v]{
						parent[v] = parent[parent[v]]	
					}

					usedVerticesLock.Lock()
						if usedVertices == graph.VerticesCount(){
							endCompressing <- struct{}{}
						}
						usedVerticesLock.Unlock()
					
				case <- endCompressing:
					endCompressing <- struct{}{}
					wgThreads.Done()
					return
				}

			}()

			wgThreads.Wait()
			usedVertices = 0
		}
	}

	//Find Representatives
	representativesSet := map[int]bool{}
	reps := []int{}
	for _, vertex := range parent{
		if !representativesSet[vertex]{
			representativesSet[vertex] = true
			reps = append(reps, vertex)
		}
	}

	return reps
}
