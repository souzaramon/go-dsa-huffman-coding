package main

import (
	"fmt"

	"github.com/souzaramon/go-dsa-huffman-coding/internal/huffman_coding"
)

func main() {
	freqMap := huffman_coding.CreateFreqMap("this is banana")
	root := huffman_coding.CreateTree(freqMap)

	fmt.Println(root)
}
