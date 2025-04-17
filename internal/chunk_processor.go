package internal

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type ChunkProcessor struct {
	InputFile     string
	TemFileDir    string
	MaxChunkLines int
}



func (cp *ChunkProcessor) ProcessChunk() ([]string, error) {
	file, err := os.Open(cp.InputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var tempfilesList []string
	var chunk []string
	chunkIndex := 0

	for scanner.Scan() {
		chunk = append(chunk, scanner.Text())

		// when the size of chunk reaches upper limits, write the chunk to temporary file
		if len(chunk) >= cp.MaxChunkLines {
			temfile, err := cp.WriteToTemFiles(chunk, chunkIndex)
			if err != nil {
				return nil, err
			}
			chunkIndex++;
			tempfilesList = append(tempfilesList, temfile)
			chunk = nil
		}
		
	}

	// process the last part
	if len(chunk) > 0{
		temfile, err := cp.WriteToTemFiles(chunk, chunkIndex)
		if err != nil {
			return nil, err
		}
		tempfilesList = append(tempfilesList, temfile)
	}

	return tempfilesList, err
}

func (cp *ChunkProcessor) WriteToTemFiles(chunk []string, index int) (string, error) {
	sort.Strings(chunk)

	tempfileName := filepath.Join(cp.TemFileDir, fmt.Sprintf("chunk-%d.tmp", index))
	file, err := os.Create(tempfileName)
	if err != nil{
		return "", err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range chunk{
		_, err := writer.WriteString(line + "\n")
		if err != nil{
			return "", err
		}
	}

	writer.Flush()
	return tempfileName, nil
}