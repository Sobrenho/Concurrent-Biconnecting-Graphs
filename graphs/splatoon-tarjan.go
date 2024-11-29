package graphs

func (graph *GraphX) SplatoonTarjan(threadsCount int) ([]int, []int) {

	components := graph.Splatoon(threadsCount)

	componentsChannel := make(chan int, len(components))
	for _, component := range components {
		componentsChannel <- component
	}

	biconnectedComponents := make([]int, len(components))

	for i := 0; i < threadsCount; i++ {

		go func() {

			u := <- componentsChannel
			if graph.Tarjan(u) {
				biconnectedComponents = append(biconnectedComponents, u)
			}

		}()

	}

	return components, biconnectedComponents
}