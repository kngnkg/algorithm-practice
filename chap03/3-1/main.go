package main

import (
	"fmt"
)

// go run main.go < input.txt

func main() {
	var v int

	var N int

	fmt.Scan(&v)
	fmt.Scan(&N)

	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}

	foundId := -1

	for i := 0; i < N; i++ {
		if a[i] == v {
			foundId = i
		}
	}

	fmt.Println(foundId)
}
