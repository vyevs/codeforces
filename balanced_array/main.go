package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1343/B
// B. Balanced Array
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// You are given a positive integer n
// , it is guaranteed that n is even (i.e. divisible by 2

// ).

// You want to construct the array a
// of length n

// such that:

//     The first n2

// elements of a are even (divisible by 2
// );
// the second n2
// elements of a are odd (not divisible by 2
// );
// all elements of a
// are distinct and positive;
// the sum of the first half equals to the sum of the second half (∑i=1n2ai=∑i=n2+1nai

//     ).

// If there are multiple answers, you can print any. It is not guaranteed that the answer exists.

// You have to answer t

// independent test cases.
// Input

// The first line of the input contains one integer t
// (1≤t≤104) — the number of test cases. Then t

// test cases follow.

// The only line of the test case contains one integer n
// (2≤n≤2⋅105) — the length of the array. It is guaranteed that that n is even (i.e. divisible by 2

// ).

// It is guaranteed that the sum of n
// over all test cases does not exceed 2⋅105 (∑n≤2⋅105

// ).
// Output

// For each test case, print the answer — "NO" (without quotes), if there is no suitable answer for the given test case or "YES" in the first line and any suitable array a1,a2,…,an
// (1≤ai≤109) satisfying conditions from the problem statement on the second line.

var bufStdin = bufio.NewReader(os.Stdin)
var bufStdout = bufio.NewWriter(os.Stdout)

func main() {
	defer bufStdout.Flush()

	var t int
	_, _ = fmt.Fscan(bufStdin, &t)

	tests := make([]int, t)
	for i := 0; i < t; i++ {
		_, _ = fmt.Fscan(bufStdin, &tests[i])
	}

	for i := 0; i < t; i++ {
		v := tests[i]

		if v%4 != 0 {
			_, _ = fmt.Fprintln(bufStdout, "NO")
		} else {
			_, _ = fmt.Fprintln(bufStdout, "YES")
			arr := constructArray(v)
			for _, v := range arr {
				_, _ = fmt.Fprint(bufStdout, v, " ")
			}
			_, _ = fmt.Fprintln(bufStdout)
		}
	}
}

func constructArray(n int) []int {
	arr := make([]int, 0, n)

	for i := 0; i < n/2; i++ {
		arr = append(arr, (i+1)*2)
	}

	for i := 0; i < n/2; i++ {
		arr = append(arr, i*2+1)
	}
	arr[n-1] += n / 2

	return arr
}
