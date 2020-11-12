package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/25/A
// A. IQ test
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Bob is preparing to pass IQ test. The most frequent task in this test is to find out which one of the given n numbers differs from the others. Bob observed that one number usually differs from the others in evenness. Help Bob — to check his answers, he needs a program that among the given n numbers finds one that is different in evenness.
// Input

// The first line contains integer n (3 ≤ n ≤ 100) — amount of numbers in the task. The second line contains n space-separated natural numbers, not exceeding 100. It is guaranteed, that exactly one of these numbers differs from the others in evenness.
// Output

// Output index of number that differs from the others in evenness. Numbers are numbered from 1 in the input order.

func main() {
	bufStdin := bufio.NewReader(os.Stdin)

	var n int
	_, _ = fmt.Fscan(bufStdin, &n)

	ns := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(bufStdin, &ns[i])
	}

	var nEvens, nOdds int
	for _, v := range ns {
		if v%2 == 0 {
			nEvens++
		} else {
			nOdds++
		}
	}

	var differingIdx int
	for i, v := range ns {
		if (nEvens == 1 && v%2 == 0) || (nOdds == 1 && v%2 == 1) {
			differingIdx = i + 1
			break
		}
	}

	_, _ = fmt.Print(differingIdx)
}
