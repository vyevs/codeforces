package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1328/A
// A. Divisibility Problem
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// You are given two positive integers a
// and b. In one move you can increase a by 1 (replace a with a+1). Your task is to find the minimum number of moves you need to do in order to make a divisible by b. It is possible, that you have to make 0 moves, as a is already divisible by b. You have to answer t

// independent test cases.
// Input

// The first line of the input contains one integer t
// (1≤t≤104) — the number of test cases. Then t

// test cases follow.

// The only line of the test case contains two integers a
// and b (1≤a,b≤109

// ).
// Output

// For each test case print the answer — the minimum number of moves you need to do in order to make a
// divisible by b.

func main() {
	bufStdin := bufio.NewReader(os.Stdin)
	bufStdout := bufio.NewWriter(os.Stdout)

	var t int
	_, _ = fmt.Fscan(bufStdin, &t)

	for i := 0; i < t; i++ {
		var a, b int
		_, _ = fmt.Fscan(bufStdin, &a, &b)

		if a <= b {
			fmt.Fprintln(bufStdout, b-a)
		} else {
			// if a is already divisible by b, we don't need take any steps
			if a%b == 0 {
				fmt.Fprintln(bufStdout, 0)
			} else {
				bsInA := a / b
				nextMultiple := (bsInA + 1) * b
				fmt.Fprintln(bufStdout, nextMultiple-a)
			}

		}
	}

	_ = bufStdout.Flush()
}
