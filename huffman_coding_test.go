package huffman_coding

import (
	"testing"
)

func TestCreateFreqMap(t *testing.T) {
	raw := "aabbcc"
	expected := map[rune]int{
		'a': 2,
		'b': 2,
		'c': 2,
	}

	freqMap := CreateFreqMap(raw)

	for symbol, count := range expected {
		if freqMap[symbol] != count {
			t.Errorf("For symbol '%c', expected %d, but got %d", symbol, count, freqMap[symbol])
		}
	}

	rawSingle := "a"
	freqMapSingle := CreateFreqMap(rawSingle)

	if freqMapSingle['a'] != 1 {
		t.Errorf("For symbol 'a', expected 1, but got %d", freqMapSingle['a'])
	}
}

func TestHuffmanComparator(t *testing.T) {
	leaf1 := HuffmanLeaf{Symbol: 'a', frequency: 5}
	leaf2 := HuffmanLeaf{Symbol: 'b', frequency: 10}

	if !HuffmanComparator(leaf1, leaf2) {
		t.Errorf("Expected 'a' to have lower frequency than 'b'")
	}

	leaf3 := HuffmanLeaf{Symbol: 'c', frequency: 10}
	if HuffmanComparator(leaf2, leaf3) {
		t.Errorf("Expected two nodes with the same frequency not to be considered smaller")
	}
}

func TestCreateTree(t *testing.T) {
	raw := "aabbcc"
	freqMap := CreateFreqMap(raw)
	tree := CreateTree(freqMap)

	if tree == nil {
		t.Error("Huffman tree not created correctly")
	}

	rawSingle := "aaaa"
	freqMapSingle := CreateFreqMap(rawSingle)
	treeSingle := CreateTree(freqMapSingle)

	if treeSingle == nil {
		t.Error("Huffman tree not created correctly for a single symbol string")
	}
}

func TestHuffmanLeafFrequency(t *testing.T) {
	leaf := HuffmanLeaf{Symbol: 'a', frequency: 5}

	if leaf.Frequency() != 5 {
		t.Errorf("Expected frequency 5, but got %d", leaf.Frequency())
	}
}

func TestHuffmanNodeFrequency(t *testing.T) {
	nodeLeft := HuffmanLeaf{Symbol: 'a', frequency: 5}
	nodeRight := HuffmanLeaf{Symbol: 'b', frequency: 10}
	node := HuffmanNode{frequency: nodeLeft.Frequency() + nodeRight.Frequency(), left: nodeLeft, right: nodeRight}

	if node.Frequency() != 15 {
		t.Errorf("Expected frequency 15, but got %d", node.Frequency())
	}
}
