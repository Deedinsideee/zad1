package main

import (
	"fmt"
)

func schet(a [3]int, n, i1, i, j int) int {

	a[i] = a[i] - a[j]
	b := a[i1]

	for g := 1; g < len(a); g++ {
		x := a[g]
		g1 := g

		for g1 > 0 && a[g1-1] > x {
			a[g1] = a[g1-1]
			g1--
		}

		a[g1] = x
	}
	if a[1] == b {
		n++

	}
	return n
}

func main() {
	var a [3]int
	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
	}

	for i1 := 0; i1 < len(a); i1++ {

		n := 0

		for i := 0; i < len(a); i++ {

			for j := 0; j < len(a); j++ {
				if i == j {
					continue
				}
				n = schet(a, n, i1, i, j)
			}
		}

		if n > 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}
