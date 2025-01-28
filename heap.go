package huffman_coding

type Heap[T any] struct {
	Data       []T
	Comparator func(a T, b T) bool
}

func (h *Heap[T]) Size() int {
	return len(h.Data)
}

func (h *Heap[T]) Peek() T {
	return h.Data[0]
}

func (h *Heap[T]) Empty() bool {
	return h.Size() == 0
}

func (h *Heap[T]) swap(a int, b int) {
	h.Data[a], h.Data[b] = h.Data[b], h.Data[a]
}

func (h *Heap[T]) bubbleUp(index int) {
	parentIndex := (index - 1) >> 1

	for index > 0 && h.Comparator(h.Data[index], h.Data[parentIndex]) {
		h.swap(index, parentIndex)
		index = parentIndex
		parentIndex = (index - 1) >> 1
	}
}

func (h *Heap[T]) sinkDown(index int) {
	childLeftIdx := index*2 + 1

	for childLeftIdx < h.Size() {
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

func (h *Heap[T]) Insert(value T) {
	h.Data = append(h.Data, value)
	h.bubbleUp(len(h.Data) - 1)
}

func (h *Heap[T]) Pop() T {
	top := h.Peek()
	last := h.Data[h.Size()-1]

	if !h.Empty() {
		h.Data[0] = last
		h.Data = h.Data[:h.Size()-1]
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
