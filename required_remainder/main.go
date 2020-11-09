package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1374/A
// A. Required Remainder
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// You are given three integers x,y
// and n. Your task is to find the maximum integer k such that 0≤k≤n that kmodx=y, where mod

// is modulo operation. Many programming languages use percent operator % to implement it.

// In other words, with given x,y
// and n you need to find the maximum possible integer from 0 to n that has the remainder y modulo x

// .

// You have to answer t
// independent test cases. It is guaranteed that such k

// exists for each test case.
// Input

// The first line of the input contains one integer t
// (1≤t≤5⋅104) — the number of test cases. The next t

// lines contain test cases.

// The only line of the test case contains three integers x,y
// and n (2≤x≤109; 0≤y<x; y≤n≤109

// ).

// It can be shown that such k

// always exists under the given constraints.
// Output

// For each test case, print the answer — maximum non-negative integer k
// such that 0≤k≤n and kmodx=y. It is guaranteed that the answer always exists.

func main() {
	bufStdin := bufio.NewReader(os.Stdin)
	bufStdout := bufio.NewWriter(os.Stdout)
	defer bufStdout.Flush()

	var t int
	_, _ = fmt.Fscan(bufStdin, &t)

	for i := 0; i < t; i++ {
		var x, y, n int
		_, _ = fmt.Fscan(bufStdin, &x, &y, &n)

		nModx := n % x

		var k int
		if nModx == y {
			k = n
		} else if nModx > y {
			k = n - (nModx - y)
		} else {
			k = n - (nModx + (x - y))
		}

		_, _ = fmt.Fprintln(bufStdout, k)
	}
}
