package internal

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
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
	var chunk []int
	chunkIndex := 0

	for scanner.Scan() {
		text := scanner.Text()
		val, err := strconv.Atoi(text)
		if err != nil {
			return nil, err
		}
		chunk = append(chunk, val)

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

func (cp *ChunkProcessor) WriteToTemFiles(chunk []int, index int) (string, error) {
	sort.Ints(chunk)

	tempfileName := filepath.Join(cp.TemFileDir, fmt.Sprintf("chunk-%d.tmp", index))
	file, err := os.Create(tempfileName)
	if err != nil{
		return "", err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range chunk{
		val := strconv.Itoa(line)
		_, err := writer.WriteString(val + "\n")
		if err != nil{
			return "", err
		}
	}

	writer.Flush()
	return tempfileName, nil
}