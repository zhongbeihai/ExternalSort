package pkg

import (
	"fmt"
	"os"
	"testing"
)

func TestExternelSort_Sort(t *testing.T) {
	tempDir := "./temp"
	os.MkdirAll(tempDir, os.ModePerm)

	sorter := &ExternalSort{
		InputFile:     "large_input.txt",
		OutputFile:    "sorted_output.txt",
		TemFileDir:     tempDir,
		MaxChunkLines: 100000, // 每个块 10 万行（你可以根据内存调整）
	}

	err := sorter.Sort()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Sorting completed successfully!")
	}
}
