package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scanWithBuffering() []int {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	var numbers []int
	for _, input := range inputs {
		val, _ := strconv.Atoi(input)
		numbers = append(numbers, val)
	}

	return numbers
}

func main() {
	var N int
	fmt.Scan(&N)

	numbers := scanWithBuffering()

	const INF = 2000000

	min := INF
	max := -INF

	for _, num := range numbers {
		if num < min {
			min = num
			continue
		}

		if num > max {
			max = num
		}
	}

	result := max - min

	fmt.Println(result)
}
