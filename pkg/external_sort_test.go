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
		InputFile:     "..\\utils\\large_input_1.txt",
		OutputFile:    "sorted_output.txt",
		TemFileDir:     tempDir,
		MaxChunkLines: 1000, // 
	}

	err := sorter.Sort()
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Sorting completed successfully!")
	}
}
