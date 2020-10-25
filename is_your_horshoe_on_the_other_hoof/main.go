package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/228/A
// A. Is your horseshoe on the other hoof?
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Valera the Horse is going to the party with friends. He has been following the fashion trends for a while, and he knows that it is very popular to wear all horseshoes of different color. Valera has got four horseshoes left from the last year, but maybe some of them have the same color. In this case he needs to go to the store and buy some few more horseshoes, not to lose face in front of his stylish comrades.

// Fortunately, the store sells horseshoes of all colors under the sun and Valera has enough money to buy any four of them. However, in order to save the money, he would like to spend as little money as possible, so you need to help Valera and determine what is the minimum number of horseshoes he needs to buy to wear four horseshoes of different colors to a party.
// Input

// The first line contains four space-separated integers s1, s2, s3, s4 (1 ≤ s1, s2, s3, s4 ≤ 109) — the colors of horseshoes Valera has.

// Consider all possible colors indexed with integers.
// Output

// Print a single integer — the minimum number of horseshoes Valera needs to buy.

func main() {
	bufStdin := bufio.NewReader(os.Stdin)

	horseshoes := make(map[int]struct{}, 4)
	var nReplaced int
	for i := 0; i < 4; i++ {
		var s int
		_, _ = fmt.Fscan(bufStdin, &s)

		if _, alreadyHaveHorseShoe := horseshoes[s]; alreadyHaveHorseShoe {
			nReplaced++
		} else {
			horseshoes[s] = struct{}{}
		}
	}

	fmt.Print(nReplaced)
}
