package graphs

import (
	"sync"
)

func (graph *Graph) SplatoonTarjan(threadsCount int) ([]int, []Block) {

    components := graph.Splatoon(threadsCount)

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