package pkg

import(
	"externalsort/pkg/internal"
)

// =======================
// controller 
// =======================

type ExternelSort struct{
	inputFile string
	outputFile string
	temFileDir string
	maxChunkLines int
}

func (es *ExternelSort) Sort() error{
	chunkProcesssor := &internal.ChunkProcessor{
		inputFile: es.inputFile,
		temFileDir: es.temFileDir,
		maxChunkLines: es.maxChunkLines,
	}

	return nil
}


