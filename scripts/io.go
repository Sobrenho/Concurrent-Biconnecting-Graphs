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
