package graphs

import (
	"encoding/binary"
	"os"
)

func (graph *Graph) WriteTo(file *os.File) error {

	err := binary.Write(file, binary.BigEndian, int64(graph.VerticesCount()))
	if err != nil {
		return err
	}

	for u := 0; u < graph.VerticesCount(); u++ {
		
		err = binary.Write(file, binary.BigEndian, int64(len(graph.Adjacents(u))))
		if err != nil {
			return err
		}

		for _, v := range graph.Adjacents(u) {

			err = binary.Write(file, binary.BigEndian, int64(v))
			if err != nil {
				return err
			}

		}

	}

	return nil
}

func ReadGraphFrom(file *os.File) (*Graph, error) {
	
	var verticesCount int64

	err := binary.Read(file, binary.BigEndian, &verticesCount)
	if err != nil {
		return nil, err
	}

	graph := NewGraph(int(verticesCount))
	
	sumOfDegrees := 0

	for u := 0; u < graph.verticesCount; u++ {

		var adjacentsSize int64

		err = binary.Read(file, binary.BigEndian, &adjacentsSize)
		if err != nil {
			return nil, err
		}

		graph.adjacents[u] = make([]int, int(adjacentsSize))

		sumOfDegrees += len(graph.adjacents[u])

		for i := range graph.adjacents[u] {

			var v int64

			err = binary.Read(file, binary.BigEndian, &v)
			if err != nil {
				return nil, err
			}

			graph.adjacents[u][i] = int(v)

		}

	}

	graph.edgesCount = sumOfDegrees / 2

	return graph, nil
}