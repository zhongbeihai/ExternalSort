package pkg

import (
	"externalsort/internal"
	"os"
)

// =======================
// controller
// =======================

type ExternalSort struct{
	InputFile string
	OutputFile string
	TemFileDir string
	MaxChunkLines int
}

func (es *ExternalSort) Sort() error{
	chunkProcesssor := &internal.ChunkProcessor{
		InputFile: es.InputFile,
		TemFileDir: es.TemFileDir,
		MaxChunkLines: es.MaxChunkLines,
	}

	temFiles, err := chunkProcesssor.ProcessChunk()
	if err != nil {
		return err
	}

	merger := &internal.Merger{
		TempFiles: temFiles,
		OutputFile: es.OutputFile,
	}
	err = merger.Merge()
	if err != nil {
		return err
	}

	for _, file := range temFiles {
		os.Remove(file)
	}

	return nil
}


