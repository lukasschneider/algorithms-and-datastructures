package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

var swaps = 0
var comps = 0

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
	swaps++

}

func downHeap(heap *[]int, pos int, last int) {
	var maximum int
	var left int
	var right int

	for pos < last {
		maximum = pos
		left = (2 * pos) + 1
		right = left + 1
		if left < last && (*heap)[left] > (*heap)[maximum] {
			maximum = left
			comps++
		}

		if right < last && (*heap)[right] > (*heap)[maximum] {
			maximum = right
			comps++
		}
		if maximum == pos {
			comps++
			return
		}
		swap(&(*heap)[pos], &(*heap)[maximum])
		pos = maximum
	}
}

func buildHeap(arr *[]int) {
	i := (len(*arr) / 2) - 1
	for i >= 0 {
		downHeap(arr, i, len(*arr))
		i--
	}
}

func heapsort(arr *[]int) {
	buildHeap(arr)
	var last int
	last = len(*arr) - 1
	for last > 0 {
		swap(&(*arr)[0], &(*arr)[last])
		downHeap(arr, 0, last)
		last--
	}
}

func main() {
	foo := []int{12, 5, 17, 6, 9, 13, 20, 3, 7, 18}
	fmt.Println(foo)
	heapsort(&foo)
	fmt.Println(foo)

	var size int
	size = 1000000
	data := make([]int, size)
	for i := range data {
		data[i] = rand.IntN(100)
	}

	swaps = 0
	comps = 0

	start := time.Now()
	heapsort(&data)
	elapsed := time.Now().Sub(start)
	fmt.Printf("Vergleiche: %d\n", comps)
	fmt.Printf("Vertauschungen: %d\n", swaps)
	fmt.Printf("Zeit: %v\n", elapsed)

}
