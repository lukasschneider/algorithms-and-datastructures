package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

var shifts int = 0

func insertionSort(arr *[]int) {
	n := len(*arr)
	for i := 1; i < n; i++ {
		j := i
		t := (*arr)[i]
		for j > 0 && (*arr)[j-1] > t {
			(*arr)[j] = (*arr)[j-1]
			j--
			shifts++
		}
		(*arr)[j] = t
	}
}

func main() {
	var size int
	size = 100000
	foo := make([]int, size)
	for i := range foo {
		foo[i] = rand.IntN(100)
	}

	data := []int{1, 25, 6, 10, 8}
	fmt.Println(data)
	insertionSort(&data)
	fmt.Println(data)
	fmt.Printf("Verschiebungen: %d\n", shifts)
	shifts = 0

	start := time.Now()
	insertionSort(&foo)
	elapsed := time.Now().Sub(start)
	fmt.Printf("Verschiebungen: %d\n", shifts)
	fmt.Printf("Zeit: %v\n", elapsed)
}
