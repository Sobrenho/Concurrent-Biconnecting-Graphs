package graphs

func (graph *GraphX) SplatoonTarjan(threadsCount int) ([]int, []int) {

	components := graph.Splatoon(threadsCount)

	componentsChannel := make(chan int, len(components))
	for _, component := range components {
		componentsChannel <- component
	}

	componentsConsumed := make(chan int, 1)
	componentsConsumed <- 0

	biconnectedComponents := make(chan []int, 1)
	biconnectedComponents <- make([]int, 0, len(components))

	canFinish := make(chan bool, 1)

	threadsFinished := make(chan int, 1)
	threadsFinished <- 0

	canReturn := make(chan bool)

	for i := 0; i < threadsCount; i++ {

		go func() {

			mustContinue := true

			for mustContinue {

				select {

				case u := <- componentsChannel:

					if graph.Tarjan(u) {
						bComponents := <- biconnectedComponents
						biconnectedComponents <- append(bComponents, u)
					}

					cConsumed := <- componentsConsumed
					cConsumed++
					componentsConsumed <- cConsumed

					if cConsumed == len(components) {
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

	return components, <- biconnectedComponents
}