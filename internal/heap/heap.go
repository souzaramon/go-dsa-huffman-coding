package heap

type Heap struct {
	Data       []int
	Comparator func(a int, b int) bool
}

func (h *Heap) swap(a int, b int) {
	h.Data[a], h.Data[b] = h.Data[b], h.Data[a]
}

func (h *Heap) size() int {
	return len(h.Data)
}

func (h *Heap) peek() int {
	return h.Data[0]
}

func (h *Heap) empty() bool {
	return h.size() == 0
}

func (h *Heap) bubbleUp(index int) {
	parentIndex := (index - 1) >> 1

	for index > 0 && h.Comparator(h.Data[index], h.Data[parentIndex]) {
		h.swap(index, parentIndex)
		index = parentIndex
		parentIndex = (index - 1) >> 1
	}
}

func (h *Heap) sinkDown(index int) {
	childLeftIdx := index*2 + 1

	for childLeftIdx < h.size() {
		childRightIdx := childLeftIdx + 1

		swapIdx := childLeftIdx

		if childRightIdx < len(h.Data) && h.Comparator(h.Data[childRightIdx], h.Data[childLeftIdx]) {
			swapIdx = childRightIdx
		}

		if h.Comparator(h.Data[swapIdx], h.Data[index]) {
			h.swap(index, swapIdx)
			index = swapIdx
			childLeftIdx = 2*index + 1
		} else {
			return
		}
	}
}

func (h *Heap) Insert(value int) {
	h.Data = append(h.Data, value)
	h.bubbleUp(len(h.Data) - 1)
}

func (h *Heap) ExtractTop() int {
	top := h.peek()
	last := h.Data[h.size()-1]

	if !h.empty() {
		h.Data[0] = last
		h.Data = h.Data[:h.size()-1]
		h.sinkDown(0)
	}

	return top
}

func MinHeapComparator(a int, b int) bool {
	return a < b
}

func MaxHeapComparator(a int, b int) bool {
	return a > b
}
