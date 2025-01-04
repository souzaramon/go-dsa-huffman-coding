package main

import (
	"fmt"

	"github.com/souzaramon/go-dsa-huffman-coding/internal/heap"
)

func main() {
	h := heap.Heap{
		Data:       make([]int, 0, 10),
		Comparator: heap.MinHeapComparator,
	}

	h.Insert(4)
	h.Insert(2)
	h.Insert(3)
	h.Insert(1)

	fmt.Println(h.Data)

	fmt.Println(h.ExtractTop())
	fmt.Println(h.Data)
}
