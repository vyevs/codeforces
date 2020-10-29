package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// https://codeforces.com/problemset/problem/1399/A
// A. Remove Smallest
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// You are given the array a
// consisting of n

// positive (greater than zero) integers.

// In one move, you can choose two indices i
// and j (i≠j) such that the absolute difference between ai and aj is no more than one (|ai−aj|≤1

// ) and remove the smallest of these two elements. If two elements are equal, you can remove any of them (but exactly one).

// Your task is to find if it is possible to obtain the array consisting of only one element using several (possibly, zero) such moves or not.

// You have to answer t

// independent test cases.
// Input

// The first line of the input contains one integer t
// (1≤t≤1000) — the number of test cases. Then t

// test cases follow.

// The first line of the test case contains one integer n
// (1≤n≤50) — the length of a. The second line of the test case contains n integers a1,a2,…,an (1≤ai≤100), where ai is the i-th element of a

// .
// Output

// For each test case, print the answer: "YES" if it is possible to obtain the array consisting of only one element using several (possibly, zero) moves described in the problem statement, or "NO" otherwise.

func main() {
	bufStdin := bufio.NewReader(os.Stdin)

	var t int
	_, _ = fmt.Fscan(bufStdin, &t)

	tests := make([][]int, t)
	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Fscan(bufStdin, &n)

		tests[i] = make([]int, n)
		for j := 0; j < n; j++ {
			_, _ = fmt.Fscan(bufStdin, &tests[i][j])
		}
	}

	for _, test := range tests {
		if canReduceTo1Elem(test) {
			_, _ = fmt.Println("YES")
		} else {
			_, _ = fmt.Println("NO")
		}

	}
}

func canReduceTo1Elem(nums []int) bool {
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] > 1 {
			return false
		}
	}

	return true
}
