package heap

import (
	"reflect"
	"testing"
)

func TestHeap_InsertMinHeap(t *testing.T) {
	h := &Heap[int]{Comparator: MinHeapComparator}
	h.Insert(3)
	h.Insert(1)
	h.Insert(4)
	h.Insert(2)

	expected := []int{1, 2, 4, 3}
	if !reflect.DeepEqual(h.Data, expected) {
		t.Errorf("Expected MinHeap to be %v, but got %v", expected, h.Data)
	}
}

func TestHeap_InsertMaxHeap(t *testing.T) {
	h := &Heap[int]{Comparator: MaxHeapComparator}
	h.Insert(3)
	h.Insert(1)
	h.Insert(4)
	h.Insert(2)

	expected := []int{4, 2, 3, 1}
	if !reflect.DeepEqual(h.Data, expected) {
		t.Errorf("Expected MaxHeap to be %v, but got %v", expected, h.Data)
	}
}

func TestHeap_ExtractTopMinHeap(t *testing.T) {
	h := &Heap[int]{Comparator: MinHeapComparator}
	h.Insert(3)
	h.Insert(1)
	h.Insert(4)
	h.Insert(2)

	top := h.Pop()

	if top != 1 {
		t.Errorf("Expected 1 when extracting top from MinHeap, but got %d", top)
	}

	expected := []int{2, 3, 4}
	if !reflect.DeepEqual(h.Data, expected) {
		t.Errorf("Expected MinHeap to be %v after extraction, but got %v", expected, h.Data)
	}
}

func TestHeap_ExtractTopMaxHeap(t *testing.T) {
	h := &Heap[int]{Comparator: MaxHeapComparator}
	h.Insert(3)
	h.Insert(1)
	h.Insert(4)
	h.Insert(2)

	top := h.Pop()

	if top != 4 {
		t.Errorf("Expected 4 when extracting top from MaxHeap, but got %d", top)
	}

	expected := []int{3, 2, 1}
	if !reflect.DeepEqual(h.Data, expected) {
		t.Errorf("Expected MaxHeap to be %v after extraction, but got %v", expected, h.Data)
	}
}

func TestHeap_EmptyHeap(t *testing.T) {
	h := &Heap[int]{Comparator: MinHeapComparator}

	if !h.Empty() {
		t.Error("Expected heap to be empty initially")
	}

	h.Insert(5)
	if h.Empty() {
		t.Error("Expected heap to not be empty after insertion")
	}
}

func TestHeap_PeekMinHeap(t *testing.T) {
	h := &Heap[int]{Comparator: MinHeapComparator}
	h.Insert(3)
	h.Insert(1)
	h.Insert(4)
	h.Insert(2)

	if h.Peek() != 1 {
		t.Errorf("Expected 1 at the top of MinHeap, but got %d", h.Peek())
	}
}

func TestHeap_PeekMaxHeap(t *testing.T) {
	h := &Heap[int]{Comparator: MaxHeapComparator}
	h.Insert(3)
	h.Insert(1)
	h.Insert(4)
	h.Insert(2)

	if h.Peek() != 4 {
		t.Errorf("Expected 4 at the top of MaxHeap, but got %d", h.Peek())
	}
}
