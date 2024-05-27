package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

var swaps int64 = 0

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
	swaps++
}

func bubblesort(arr *[]int) {
	for n := len(*arr) - 1; n > 1; n-- {
		for i := 0; i < n; i++ {
			if (*arr)[i] > (*arr)[i+1] {
				swap(&(*arr)[i], &(*arr)[i+1])
			}
		}
	}
}

func main() {

	var size int
	size = 1000000
	foo := make([]int, size)
	for i := range foo {
		foo[i] = rand.IntN(100)
	}

	data := []int{1, 25, 6, 10, 8}
	fmt.Println(data)
	bubblesort(&data)
	fmt.Println(data)
	fmt.Printf("Vertauschungen: %d\n", swaps)
	swaps = 0

	bubblesort(&foo)
	start := time.Now()
	elapsed := time.Now().Sub(start)
	fmt.Printf("Vergleiche: %d\n", swaps)
	fmt.Printf("Zeit: %v\n", elapsed)
}
