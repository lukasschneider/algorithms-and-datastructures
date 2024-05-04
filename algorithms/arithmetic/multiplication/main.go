package main

import "fmt"

func main() {
	var x, y, z int
	a := 10 // Setze a entsprechend deiner Anforderungen
	b := 7  // Setze b entsprechend deiner Anforderungen

	x = a
	y = b
	z = 0

	for y > 0 {
		if y%2 != 0 {
			y = y - 1
			z = z + x
		} else {
			y = y / 2
			x = x + x
		}
	}

	fmt.Println("Ergebnis:", z)
}
