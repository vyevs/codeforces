package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// https://codeforces.com/problemset/problem/492/B
// B. Vanya and Lanterns
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Vanya walks late at night along a straight street of length l, lit by n lanterns. Consider the coordinate system with the beginning of the street corresponding to the point 0, and its end corresponding to the point l. Then the i-th lantern is at the point ai. The lantern lights all points of the street that are at the distance of at most d from it, where d is some positive number, common for all lanterns.

// Vanya wonders: what is the minimum light radius d should the lanterns have to light the whole street?
// Input

// The first line contains two integers n, l (1 ≤ n ≤ 1000, 1 ≤ l ≤ 109) — the number of lanterns and the length of the street respectively.

// The next line contains n integers ai (0 ≤ ai ≤ l). Multiple lanterns can be located at the same point. The lanterns may be located at the ends of the street.
// Output

// Print the minimum light radius d, needed to light the whole street. The answer will be considered correct if its absolute or relative error doesn't exceed 10 - 9.

func main() {
	bufStdin := bufio.NewReader(os.Stdin)

	var n, l int
	_, _ = fmt.Fscan(bufStdin, &n, &l)

	as := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(bufStdin, &as[i])
	}

	sort.Ints(as)

	minDist := float64(as[0])
	for i := 0; i < n-1; i++ {
		l1 := as[i]
		l2 := as[i+1]

		dist := float64(l2-l1) / 2.0

		if dist > minDist {
			minDist = dist
		}
	}

	lastDist := float64(l - as[len(as)-1])
	if lastDist > minDist {
		minDist = lastDist
	}

	_, _ = fmt.Print(minDist)
}
