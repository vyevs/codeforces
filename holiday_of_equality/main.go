package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/758/A
// A. Holiday Of Equality
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// In Berland it is the holiday of equality. In honor of the holiday the king decided to equalize the welfare of all citizens in Berland by the expense of the state treasury.

// Totally in Berland there are n citizens, the welfare of each of them is estimated as the integer in ai burles (burle is the currency in Berland).

// You are the royal treasurer, which needs to count the minimum charges of the kingdom on the king's present. The king can only give money, he hasn't a power to take away them.
// Input

// The first line contains the integer n (1 ≤ n ≤ 100) — the number of citizens in the kingdom.

// The second line contains n integers a1, a2, ..., an, where ai (0 ≤ ai ≤ 106) — the welfare of the i-th citizen.
// Output

// In the only line print the integer S — the minimum number of burles which are had to spend.

func main() {
	bufStdin := bufio.NewReader(os.Stdin)

	var n int
	_, _ = fmt.Fscan(bufStdin, &n)

	as := make([]int, n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(bufStdin, &as[i])
	}

	m := max(as)

	var cost int
	for _, v := range as {
		cost += m - v
	}
	_, _ = fmt.Print(cost)
}

func max(s []int) int {
	m := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > m {
			m = s[i]
		}
	}
	return m
}
