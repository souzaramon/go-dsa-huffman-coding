package huffman_coding

import (
	"sort"
)

type Node interface {
	Frequency() int
}

type HuffmanLeaf struct {
	Symbol    rune
	frequency int
}

func (hf HuffmanLeaf) Frequency() int {
	return hf.frequency
}

type HuffmanNode struct {
	frequency   int
	left, right Node
}

func (hn HuffmanNode) Frequency() int {
	return hn.frequency
}

func HuffmanComparator(a Node, b Node) bool {
	return a.Frequency() < b.Frequency()
}

func CreateFreqMap(raw string) map[rune]int {
	freqMap := map[rune]int{}

	for _, symbol := range raw {
		acc, exists := freqMap[symbol]

		if exists {
			freqMap[symbol] = acc + 1
		} else {
			freqMap[symbol] = 1
		}
	}

	return freqMap
}

func CreateTree(freqMap map[rune]int) Node {
	tree := Heap[Node]{
		Data:       make([]Node, 0),
		Comparator: HuffmanComparator,
	}

	var freqMapkeys []rune

	for k := range freqMap {
		freqMapkeys = append(freqMapkeys, k)
	}

	sort.Slice(freqMapkeys, func(i, j int) bool {
		return freqMapkeys[i] < freqMapkeys[j]
	})

	for _, key := range freqMapkeys {
		tree.Insert(HuffmanLeaf{Symbol: key, frequency: freqMap[key]})
	}

	for tree.Size() > 1 {
		a := tree.Pop()
		b := tree.Pop()

		tree.Insert(HuffmanNode{frequency: a.Frequency() + b.Frequency(), left: a, right: b})
	}

	return tree.Pop()
}

func CreateCodes() {

}
