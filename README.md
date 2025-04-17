# 🧊 Go External Sort

A scalable, concurrent external sorting program written in Go, capable of sorting files larger than system memory.

---

## 📌 Project Overview

This project implements a **two-phase external sort algorithm** in Go:

1. **Chunking Phase (Split & Sort)**  
   The large input file is split into smaller, manageable chunks which are sorted individually in-memory and written to temporary files (sorted chunks).

2. **Merge Phase (Multi-way Merge)**  
   All sorted chunk files are merged into a single fully sorted output file using a min-heap-based k-way merge algorithm.

> This solution is designed to sort files of 10GB, 100GB or more — much larger than RAM capacity.

---

## 🛠 Technical Architecture

### 🔹 Main Components

| Component        | Description |
|------------------|-------------|
| `ExternalSorter` | Main controller that coordinates the sort process. |
| `ChunkProcessor` | Responsible for reading, sorting, and writing file chunks concurrently. |
| `Merger`         | Performs k-way merge on the sorted chunk files. |
| `MinHeap`        | A priority queue (min-heap) used for efficient merging. |

---

## 🚀 Concurrency Model

- **ChunkProcessor** uses a producer-consumer model:
  - The main goroutine reads chunks from the input file.
  - A fixed number of worker goroutines (`workerCount`) sort and write those chunks in parallel.
  
- This dramatically improves performance on multi-core systems.

---

## 🧠 Algorithm Details

### ✅ Phase 1: Chunking (Split + Sort)

- Input file is scanned line by line.
- Lines are grouped into chunks of `maxChunkLines`.
- Each chunk is sorted using `sort.Strings()` (can be replaced by custom sort logic).
- Sorted data is written to a temp file: `chunk_<workerID>_<timestamp>.tmp`.

### ✅ Phase 2: Multi-way Merge

- All sorted chunk files are opened with buffered readers.
- The first line from each file is inserted into a **min-heap**.
- The smallest element is repeatedly popped and written to the output file.
- After each pop, the next element from the same source file is pushed to the heap.

---

## 📂 Project Structure

