package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1335/A
// A. Candies and Two Sisters
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// There are two sisters Alice and Betty. You have n
// candies. You want to distribute these n

// candies between two sisters in such a way that:

//     Alice will get a

// (a>0
// ) candies;
// Betty will get b
// (b>0
// ) candies;
// each sister will get some integer number of candies;
// Alice will get a greater amount of candies than Betty (i.e. a>b
// );
// all the candies will be given to one of two sisters (i.e. a+b=n

//     ).

// Your task is to calculate the number of ways to distribute exactly n

// candies between sisters in a way described above. Candies are indistinguishable.

// Formally, find the number of ways to represent n
// as the sum of n=a+b, where a and b are positive integers and a>b

// .

// You have to answer t

// independent test cases.
// Input

// The first line of the input contains one integer t
// (1≤t≤104) — the number of test cases. Then t

// test cases follow.

// The only line of a test case contains one integer n
// (1≤n≤2⋅109

// ) — the number of candies you have.
// Output

// For each test case, print the answer — the number of ways to distribute exactly n
// candies between two sisters in a way described in the problem statement. If there is no way to satisfy all the conditions, print 0.

func main() {
	bufStdin := bufio.NewReader(os.Stdin)
	bufStdout := bufio.NewWriter(os.Stdout)

	var t int
	_, _ = fmt.Fscan(bufStdin, &t)

	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Fscan(bufStdin, &n)

		_, _ = fmt.Fprintln(bufStdout, nDistributions(n))
	}

	_ = bufStdout.Flush()
}

// for an even number, there are (n/2)-1 ways to distribute the candies
// e.g. 3 ways to divide 8: 5-3, 6-2, 7-1
// exception is 2, for which there are no distributions
//
// for an odd number, there are ((n+1)/2)-1 ways to distribute the candies
// e.g 3 ways to divide 7: 4-3, 5-2, 6-1
// exception is 1, for which there are no distributions
//
// due to integer division, ((n+1)/2)-1 is correct for both even and odd numbers
func nDistributions(n int) int {
	return ((n + 1) / 2) - 1
}
