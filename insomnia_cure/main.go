package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/148/A
// A. Insomnia cure
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// «One dragon. Two dragon. Three dragon», — the princess was counting. She had trouble falling asleep, and she got bored of counting lambs when she was nine.

// However, just counting dragons was boring as well, so she entertained herself at best she could. Tonight she imagined that all dragons were here to steal her, and she was fighting them off. Every k-th dragon got punched in the face with a frying pan. Every l-th dragon got his tail shut into the balcony door. Every m-th dragon got his paws trampled with sharp heels. Finally, she threatened every n-th dragon to call her mom, and he withdrew in panic.

// How many imaginary dragons suffered moral or physical damage tonight, if the princess counted a total of d dragons?
// Input

// Input data contains integer numbers k, l, m, n and d, each number in a separate line (1 ≤ k, l, m, n ≤ 10, 1 ≤ d ≤ 105).
// Output

// Output the number of damaged dragons.

func main() {
	bufStdin := bufio.NewReader(os.Stdin)

	var k, l, m, n, d int
	_, _ = fmt.Fscan(bufStdin, &k, &l, &m, &n, &d)

	var nHurt int
	for i := 1; i <= d; i++ {
		if (i%k == 0) || (i%l == 0) || (i%m == 0) || (i%n == 0) {
			nHurt++
		}
	}
	fmt.Print(nHurt)
}
