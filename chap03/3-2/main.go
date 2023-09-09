package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var v int

	var N int

	fmt.Scan(&v)
	fmt.Scan(&N)

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	var a []int
	for _, input := range inputs {
		val, _ := strconv.Atoi(input)
		a = append(a, val)
	}

	count := 0
	for i := 0; i < N; i++ {
		if a[i] == v {
			count++
		}
	}

	fmt.Println(count)
}
