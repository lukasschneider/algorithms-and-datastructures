package selection

import (
	"fmt"
	"math/rand/v2"
	"time"
)

var ifs int = 0

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}

func selectionSort(arr *[]int) {
	var minimum int
	n := len(*arr)
	for i := 0; i < n-1; i++ {
		minimum = i
		for j := i + 1; j < n; j++ {
			ifs++
			if (*arr)[j] < (*arr)[minimum] {
				minimum = j
			}
		}
		swap(&(*arr)[i], &(*arr)[minimum])
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
	selectionSort(&data)
	fmt.Println(data)
	fmt.Printf("Vergleiche: %d\n", ifs)
	ifs = 0

	start := time.Now()
	selectionSort(&foo)
	elapsed := time.Now().Sub(start)
	fmt.Printf("Vergleiche: %d\n", ifs)
	fmt.Printf("Zeit: %v\n", elapsed)
}
