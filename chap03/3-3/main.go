package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const INF = 100000

func main() {
	var N int

	fmt.Scan(&N)

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	var numbers []int
	for _, input := range inputs {
		val, _ := strconv.Atoi(input)
		numbers = append(numbers, val)
	}

	var min [2]int = [2]int{INF, INF}
	for _, num := range numbers {
		if num < min[0] {
			min[1] = min[0]
			min[0] = num
			continue
		}

		if num < min[1] {
			min[1] = num
		}
	}

	fmt.Println(min[1])
}
