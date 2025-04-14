package internel


type HeapItem struct {
	value     string
	fileIndex int
}

type MinHeap []HeapItem

func (h MinHeap) Len() int               { return len(h) }
func (h MinHeap) Less(i int, j int) bool { return h[i].value < h[j].value }
func (h MinHeap) Swap(i int, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{})    { *h = append(*h, x.(HeapItem)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n - 1]
	*h = old[0: n - 1]
	return item
}
