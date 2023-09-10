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

func myAnswer() {
	numbers := scanWithBuffering()

	// 配列が全て偶数であるかを調べる
	isAllEven := func(numbers []int) bool {
		for _, num := range numbers {
			if num%2 != 0 {
				return false
			}
		}

		return true
	}

	count := 0

	for {
		if !isAllEven(numbers) {
			break
		}

		for i, num := range numbers {
			numbers[i] = num / 2
		}

		count++
	}

	fmt.Println(count)
}

func modelAnswer() {
	numbers := scanWithBuffering()

	// Nが2で何回割れるか
	howManyTimes := func(N int) int {
		exp := 0
		for {
			if N%2 != 0 {
				break
			}

			N = N / 2
			exp++
		}

		return exp
	}

	const INF = 100000

	result := INF
	for _, num := range numbers {
		// 2で割り切れる回数が最も少ないものを調べる
		result = min(result, howManyTimes(num))
	}

	fmt.Println(result)
}

func main() {
	// myAnswer()
	modelAnswer()
}
