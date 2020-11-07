package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1399/B
// B. Gifts Fixing
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// You have n
// gifts and you want to give all of them to children. Of course, you don't want to offend anyone, so all gifts should be equal between each other. The i-th gift consists of ai candies and bi

// oranges.

// During one move, you can choose some gift 1≤i≤n

// and do one of the following operations:

//     eat exactly one candy from this gift (decrease ai

// by one);
// eat exactly one orange from this gift (decrease bi
// by one);
// eat exactly one candy and exactly one orange from this gift (decrease both ai
// and bi

//     by one).

// Of course, you can not eat a candy or orange if it's not present in the gift (so neither ai
// nor bi

// can become less than zero).

// As said above, all gifts should be equal. This means that after some sequence of moves the following two conditions should be satisfied: a1=a2=⋯=an
// and b1=b2=⋯=bn (and ai equals bi

// is not necessary).

// Your task is to find the minimum number of moves required to equalize all the given gifts.

// You have to answer t

// independent test cases.
// Input

// The first line of the input contains one integer t
// (1≤t≤1000) — the number of test cases. Then t

// test cases follow.

// The first line of the test case contains one integer n
// (1≤n≤50) — the number of gifts. The second line of the test case contains n integers a1,a2,…,an (1≤ai≤109), where ai is the number of candies in the i-th gift. The third line of the test case contains n integers b1,b2,…,bn (1≤bi≤109), where bi is the number of oranges in the i

// -th gift.
// Output

// For each test case, print one integer: the minimum number of moves required to equalize all the given gifts.

func main() {
	bufStdin := bufio.NewReader(os.Stdin)
	bufStdout := bufio.NewWriter(os.Stdout)
	defer bufStdout.Flush()

	var t int
	_, _ = fmt.Fscan(bufStdin, &t)

	for i := 0; i < t; i++ {
		var i uint64
		var n uint64
		_, _ = fmt.Fscan(bufStdin, &n)

		as := make([]uint64, n)
		for i = 0; i < n; i++ {
			_, _ = fmt.Fscan(bufStdin, &as[i])
		}

		bs := make([]uint64, n)
		for i = 0; i < n; i++ {
			_, _ = fmt.Fscan(bufStdin, &bs[i])
		}

		// we must reduce the values in as to the min of the values, e.g. [4, 5, 6] -> [4, 4, 4], this is 3 moves
		// same for the bs array, these are the targets we want to reach for the 2 arrays
		candyTarget := min(as)
		orangeTarget := min(bs)

		var nSteps uint64
		for i = 0; i < n; i++ {
			// for a gift to reach the wanted state, both the candies and oranges must equal their intended target
			// and the minmum number of steps we need is the max steps between the 2, and the other's steps will be done during
			// a subset of the max steps
			nSteps += max([]uint64{as[i] - candyTarget, bs[i] - orangeTarget})
		}

		_, _ = fmt.Fprintln(bufStdout, nSteps)
	}

}

func min(vs []uint64) uint64 {
	m := vs[0]
	for i := 1; i < len(vs); i++ {
		if vs[i] < m {
			m = vs[i]
		}
	}
	return m
}

func max(vs []uint64) uint64 {
	m := vs[0]
	for i := 1; i < len(vs); i++ {
		if vs[i] > m {
			m = vs[i]
		}
	}
	return m
}
