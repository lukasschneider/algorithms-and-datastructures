package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

var swaps = 0

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
	swaps++

}

func partition(arr *[]int, left int, right int) int {
	pivot := (*arr)[left]
	i := left + 1
	j := right
	for {
		for i < right && (*arr)[i] <= pivot {
			i++
		}
		for j > left && (*arr)[j] >= pivot {
			j--
		}
		if i < j {
			swap(&(*arr)[i], &(*arr)[j])
		}
		if j <= i {
			break
		}
	}
	swap(&(*arr)[left], &(*arr)[j])
	return j
}

func quicksort(arr *[]int, left int, right int) {
	if left < right {
		m := partition(&(*arr), left, right)
		quicksort(arr, left, m-1)
		quicksort(arr, m+1, right)
	}
}

func main() {
	var size int
	size = 100000
	foo := make([]int, size)
	for i, _ := range foo {
		foo[i] = rand.IntN(100)
	}

	data := []int{1, 25, 6, 10, 8}
	fmt.Println(data)
	quicksort(&data, 0, len(data)-1)
	fmt.Println(data)
	fmt.Printf("Vertauschungen: %d\n", swaps)
	swaps = 0

	quicksort(&foo, 0, len(foo)-1)
	start := time.Now()
	elapsed := time.Now().Sub(start)
	fmt.Printf("Vergleiche: %d\n", swaps)
	fmt.Printf("Zeit: %v\n", elapsed)
}
