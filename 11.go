package main

import "fmt"

func sets() {
	a := []int{5, 39, 82, 101, 99, 52, 1, 37, 21, 11}
	b := []int{3, 221, 15, 82, 5, 16, 2, 21, 11, 9, 39}

	if len(a) < len(b) {
		a, b = b, a
	}

	result := []int{}
	for _, y := range a {
		for _, j := range b {
			if j == y {
				result = append(result, j)
			}
		}
	}

	fmt.Println(result)
}
