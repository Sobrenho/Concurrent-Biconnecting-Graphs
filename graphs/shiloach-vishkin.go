package graphs

import (
	"sync"
)


func (graph *Graph) ShiloachVishkin(threadsCount int) []int{


	parent := make([]int, graph.VerticesCount())

	vertChanHooking := make(chan int, graph.VerticesCount())
	vertChanCompressing := make(chan int, graph.VerticesCount())

	endHooking := make(chan bool, 1) //Faz o papel de mutex
	endCompressing := make(chan bool, 1) //Faz o papel de mutex

	for i := 0; i < graph.VerticesCount(); i++ {
		parent[i] = i
		vertChanHooking <- i
	}
	
	var updateLock sync.Mutex

	var waitGroupBarrier sync.WaitGroup
	waitGroupBarrier.Add(0)

	var waitGroupThreads sync.WaitGroup
	waitGroupThreads.Add(threadsCount)

	update := true

	for i := 0; i< threadsCount; i++{
		go func() {
			for{	
				updateLock.Lock()
				update = false
				updateLock.Unlock()

				//Hooking
				select{
				case v := <- vertChanHooking:
					cur_parent := parent[v]
					for _, neighbor := range graph.Adjacents(v){
						if parent[neighbor] < cur_parent{

							parent[v] = parent[neighbor]

							updateLock.Lock()
							update = true
							updateLock.Unlock()

						}

					}

					if v == graph.VerticesCount(){
						<- endHooking
						endHooking <- true
						break
					}

					vertChanCompressing <- v

				case <- endHooking: //Talvez esteja sobrando um endHooking sem 
					endHooking <- true
					break	
				}

				//Barreira


				//Compressing
				select{

				case v := <- vertChanCompressing:
					for parent[parent[v]] != parent[v]{
							
						parent[v] = parent[parent[v]]
						
					}
				
				case <- endCompressing:
					endCompressing <- true
					break
				}
				
				
				
				
				//Verificar se teve mudança ou podemos parar
				updateLock.Lock()
				if update == false{
					updateLock.Unlock()
					waitGroupThreads.Done()
					return
				}
				updateLock.Unlock()

				//Barreira
			}
		}()
	}
	
	waitGroupThreads.Wait()
	

	//Find Representatives
	representativesSet := map[int]bool{}
	unique := []int{}
	for _, vertex := range parent{
		if !representativesSet[vertex]{
			representativesSet[vertex] = true
			unique = append(unique, vertex)
		}
	}

	return unique
}


func (graph *Graph) ShiloachVishkinTarjan(threadsCount int)([]int, []Block){

	components := graph.ShiloachVishkin(threadsCount)


	nextComponent := 0
    var nextComponentLock sync.Mutex

    blocks := make([]Block, 0)
    var blocksLock sync.Mutex

    var waitGroup sync.WaitGroup

    waitGroup.Add(threadsCount)
    for i := 0; i < threadsCount; i++ {

        go func() {

            for {

                nextComponentLock.Lock()
                if nextComponent == len(components) {
                    nextComponentLock.Unlock()
                    break
                }
                component := components[nextComponent]
                nextComponent++
                nextComponentLock.Unlock()

                blocksHere := graph.Tarjan(component)

                blocksLock.Lock()
                blocks = append(blocks, blocksHere...)
                blocksLock.Unlock()
            }

            waitGroup.Done()
        }()

    }

    waitGroup.Wait()
    return components, blocks
}