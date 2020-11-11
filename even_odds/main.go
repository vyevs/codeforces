package main

import "fmt"

// https://codeforces.com/problemset/problem/318/A
// A. Even Odds
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Being a nonconformist, Volodya is displeased with the current state of things, particularly with the order of natural numbers (natural number is positive integer number). He is determined to rearrange them. But there are too many natural numbers, so Volodya decided to start with the first n. He writes down the following sequence of numbers: firstly all odd integers from 1 to n (in ascending order), then all even integers from 1 to n (also in ascending order). Help our hero to find out which number will stand at the position number k.
// Input

// The only line of input contains integers n and k (1 ≤ k ≤ n ≤ 1012).

// Please, do not use the %lld specifier to read or write 64-bit integers in C++. It is preferred to use the cin, cout streams or the %I64d specifier.
// Output

// Print the number that will stand at the position number k after Volodya's manipulations.

func main() {
	var n, k uint64
	_, _ = fmt.Scan(&n, &k)

	var v uint64
	if k > (n+1)/2 {
		v = (k - (n+1)/2) * 2
	} else {
		v = k*2 - 1
	}

	_, _ = fmt.Print(v)
}
