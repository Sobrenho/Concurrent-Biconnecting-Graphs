package scripts

import (
	"encoding/binary"
	"os"
	"splatoon-tarjan-demo/graphs"
)

func writeEdgeToFile(edge graphs.Edge, file *os.File) error {

	err := binary.Write(file, binary.BigEndian, int64(edge.U))
	if err != nil {
		return err
	}

	err = binary.Write(file, binary.BigEndian, int64(edge.V))
	if err != nil {
		return err
	}

	return nil
}

func readEdgeFromFile(edge *graphs.Edge, file *os.File) error {
	
	var u int64
	err := binary.Read(file, binary.BigEndian, &u)
	if err != nil {
		return err
	}

	var v int64
	err = binary.Read(file, binary.BigEndian, &v)
	if err != nil {
		return err
	}

	edge.U = int(u)
	edge.V = int(v)

	return nil
}

func writeBlockToFile(block graphs.Block, file *os.File) error {

	err := binary.Write(file, binary.BigEndian, int64(len(block)))
	if err != nil {
		return err
	}

	for _, item := range block {
		err = writeEdgeToFile(item, file)
		if err != nil {
			return err
		}
	}

	return nil
}

func readSliceOfEdgesFromFile(slice *[]graphs.Edge, file *os.File) error {

	var size int64
	err := binary.Read(file, binary.BigEndian, &size)
	if err != nil {
		return err
	}

	*slice = make([]graphs.Edge, int(size))

	for i := range *slice {
		err = readEdgeFromFile(&(*slice)[i], file)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeSliceOfBlocksToFile(slice []graphs.Block, file *os.File) error {

	err := binary.Write(file, binary.BigEndian, int64(len(slice)))
	if err != nil {
		return err
	}

	for _, subslice := range slice {
		err = writeBlockToFile(subslice, file)
		if err != nil {
			return err
		}
	}

	return nil
}

func readSliceOfSlicesOfEdgesFromFile(slice *[][]graphs.Edge, file *os.File) error {

	var size int64
	err := binary.Read(file, binary.BigEndian, &size)
	if err != nil {
		return err
	}

	*slice = make([][]graphs.Edge, int(size))

	for i := range *slice {
		err = readSliceOfEdgesFromFile(&(*slice)[i], file)
		if err != nil {
			return err
		}
	}

	return nil
}
