package internal

import (
	"bufio"
	"container/heap"
	"os"
)

type Merger struct {
	TempFiles  []string
	OutputFile string
}

func (m *Merger) Merge() error {
	// assign a Reader to every tempFile
	readers := make([]*bufio.Scanner, len(m.TempFiles))
	files := make([]*os.File, len(m.TempFiles))

	for i, temFile := range m.TempFiles{
		file, err := os.Open(temFile)
		if err != nil{
			return err
		}
		files[i] = file
		readers[i] = bufio.NewScanner(file)
	}

	defer func ()  {
		for _ , f := range files{
			f.Close()
		}
	}()

	// Initialize Heap
	h := &MinHeap{}
	heap.Init(h)

	for i, reader := range readers{
		if reader.Scan(){
			heap.Push(h, HeapItem{value: reader.Text(), fileIndex: i})
		}
	}

	output, err := os.Create(m.OutputFile)
	if  err != nil {
		return err
	}
	defer output.Close()

	writer := bufio.NewWriter(output)

	for h.Len() > 0{
		item := heap.Pop(h).(HeapItem)

		_, err := writer.WriteString(item.value + "\n")
		if err != nil {
			return err
		}

		if readers[item.fileIndex].Scan() {
			heap.Push(h, HeapItem{value: readers[item.fileIndex].Text(), fileIndex: item.fileIndex})
		}
	}

	writer.Flush()
	return nil
}