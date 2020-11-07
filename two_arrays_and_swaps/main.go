package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// https://codeforces.com/problemset/problem/1353/B
// B. Two Arrays And Swaps
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// You are given two arrays a
// and b both consisting of n positive (greater than zero) integers. You are also given an integer k

// .

// In one move, you can choose two indices i
// and j (1≤i,j≤n) and swap ai and bj (i.e. ai becomes bj and vice versa). Note that i and j can be equal or different (in particular, swap a2 with b2 or swap a3 and b9

// both are acceptable moves).

// Your task is to find the maximum possible sum you can obtain in the array a
// if you can do no more than (i.e. at most) k

// such moves (swaps).

// You have to answer t

// independent test cases.
// Input

// The first line of the input contains one integer t
// (1≤t≤200) — the number of test cases. Then t

// test cases follow.

// The first line of the test case contains two integers n
// and k (1≤n≤30;0≤k≤n) — the number of elements in a and b and the maximum number of moves you can do. The second line of the test case contains n integers a1,a2,…,an (1≤ai≤30), where ai is the i-th element of a. The third line of the test case contains n integers b1,b2,…,bn (1≤bi≤30), where bi is the i-th element of b

// .
// Output

// For each test case, print the answer — the maximum possible sum you can obtain in the array a
// if you can do no more than (i.e. at most) k swaps.

func main() {
	bufStdin := bufio.NewReader(os.Stdin)
	bufStdout := bufio.NewWriter(os.Stdout)
	defer bufStdout.Flush()

	var t int
	_, _ = fmt.Fscan(bufStdin, &t)

	for i := 0; i < t; i++ {
		var n, k int
		_, _ = fmt.Fscan(bufStdin, &n, &k)

		a := make([]int, n)
		for i := 0; i < n; i++ {
			_, _ = fmt.Fscan(bufStdin, &a[i])
		}

		b := make([]int, n)
		for i := 0; i < n; i++ {
			_, _ = fmt.Fscan(bufStdin, &b[i])
		}

		sort.Ints(a)
		sort.Ints(b)

		// swap the minimum values of a with the maximum values of b
		// the only special case is that we should not swap if the b value is less than a, at that point we can stop looping
		// since we have reached a point where the b values will not be greater in the future because the slices are sorted
		for i := 0; i < k && b[n-i-1] > a[i]; i++ {
			a[i] = b[n-i-1]
		}

		var sum int
		for i := 0; i < n; i++ {
			sum += a[i]
		}
		_, _ = fmt.Fprintln(bufStdout, sum)
	}
}
