package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

var shifts int = 0

func shellSort(arr *[]int) {
	var j int
	n := len(*arr)
	for interval := n / 2; interval > 0; interval /= 2 {
		for i := interval; i < n; i++ {
			t := (*arr)[i]
			for j = i; j >= interval && (*arr)[j-interval] > t; j -= interval {
				(*arr)[j] = (*arr)[j-interval]
				shifts++
			}
			(*arr)[j] = t
		}
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
	shellSort(&data)
	fmt.Println(data)
	fmt.Printf("Verschiebungen: %d\n", shifts)
	shifts = 0

	start := time.Now()
	shellSort(&foo)
	elapsed := time.Now().Sub(start)
	fmt.Printf("Verschiebungen: %d\n", shifts)
	fmt.Printf("Zeit: %v\n", elapsed)
}
