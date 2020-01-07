package main

import (
	"fmt"
)

func main() {
	slc := generateSortedData(100)
	fmt.Println(binarySearch(slc, 50, 0))

	fmt.Println(binarySearchFloat([]float64{0.02, 0.2, 0.23, 0.3}, 0.02, 0))
}

func generateSortedData(len int) []int {
	slc := make([]int, len)
	for i := 0; i < len; i++ {
		slc[i] = i * 2
	}
	return slc
}

func binarySearch(slc []int, searchElem, offset int) int {
	if len(slc) == 0 {
		return -1
	}
	if len(slc) == 1 {
		if slc[0] == searchElem {
			return offset
		}
		return -1
	}
	middle := len(slc) >> 1
	if slc[middle] > searchElem {
		return binarySearch(slc[:middle-1], searchElem, offset-middle)
	}
	return binarySearch(slc[middle:], searchElem, offset+middle)
}

func binarySearchFloat(slc []float64, searchElem float64, offset int) int {
	if len(slc) == 0 {
		return -1
	}
	if len(slc) == 1 {
		if slc[0] == searchElem {
			return offset
		}
		return -1
	}

	middle := len(slc) /2

	fmt.Println( len(slc) , searchElem)
	if slc[middle] > searchElem {
		fmt.Println("hui", offset, "-", middle)
		return binarySearchFloat(slc[:middle], searchElem, offset-middle)
	}
	fmt.Println("hui", offset, "+", middle)
	return binarySearchFloat(slc[middle:], searchElem, offset+middle)
}
	