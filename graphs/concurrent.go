package graphs

import (
	"runtime"
)



func (G *Graph) BiconnectedComponentsInGraphConcurrentCheck(max_threads int){
	runtime.GOMAXPROCS(max_threads)

	vertexForEachComponents := G.SplatoonComponentSearch(max_threads)
	for _, vertex := range vertexForEachComponents{
		go G.CheckIfBiconnectedComponentWithSourceVertex(vertex)
	}
}