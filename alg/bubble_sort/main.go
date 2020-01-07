package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slc := generateSlice(13)
	fmt.Println(bubbleSort(slc))
}

func bubbleSort(slc []int) []int {
	step := 0

	for i := 0; i < len(slc); i++ {

		for j := 0; j < len(slc)-step; j++ {
			if len(slc)-1 == j {
				continue
			}

			if slc[j] > slc[j+1] {
				slc[j], slc[j+1] = slc[j+1], slc[j]
			}
		}

		step++
	}
	return slc
}

func generateSlice(size int) []int {
	slc := make([]int, size, size)
	for i := range slc {
		slc[i] = rand.Intn(999)

	}
	return slc
}
