package main

import "fmt"

func main() {
	var K int
	var N int
	fmt.Scan(&K)
	fmt.Scan(&N)

	count := 0
	for x := 0; x <= min(K, N); x++ {
		for y := 0; y <= min(K, N); y++ {
			// x + y + z = N は
			// z = N - x - y と変形すれば、zの値がただ1つに決まる
			z := N - x - y

			// zが条件を満たしているか確認
			if z >= 0 && z <= K {
				count++
			}
		}
	}

	fmt.Println(count)
}
