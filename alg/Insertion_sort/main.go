package main

import (
	"fmt"
	"math/rand"

	"github.com/apaxa-go/helper/mathh"
)

func main() {
	slc := generateSlice(10)

	fmt.Println(insertionSort(slc))
}

func generateSlice(size int) []int {
	slc := make([]int, size, size)
	for i := range slc {
		slc[i] = rand.Intn(mathh.MaxInt)

	}
	return slc
}

func insertionSort(items []int) []int {
	for i := 1; i < len(items); i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
	return items
}
