package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/486/A
// A. Calculating Function
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// For a positive integer n let's define a function f:

// f(n) =  - 1 + 2 - 3 + .. + ( - 1)nn

// Your task is to calculate f(n) for a given integer n.
// Input

// The single line contains the positive integer n (1 ≤ n ≤ 10^15).
// Output

// Print f(n) in a single line.

func main() {
	solveTheProblem()
}

func solveTheProblem() {
	bufStdin := bufio.NewReader(os.Stdin)

	// the range of n can go as far as 10^15 so we need 64 bits
	var n int64
	fmt.Fscan(bufStdin, &n)

	// right away, the possible range of n tells us that we should not actually do the n sums to determine the result
	// 10^15 is much too large, we need a more clever way to calculate the result
	// we can find a pattern:
	// f(1) = -1
	// f(2) = 1
	// f(3) = -2
	// f(4) = 2
	// f(5) = -3
	// f(6) = 3
	// etc...
	// we can see that if n is even, f(n) = n/2 and if n is odd, then f(n) = -((n+1)/2)
	// due to the way integer division works e.g. 7/2 = 3, we can condense the magnitude of the result into (n+1)/2
	// the sign is negative is n%2==1 and positive if n%2==0

	var sign int64
	if n%2 == 0 {
		sign = 1
	} else {
		sign = -1
	}
	mag := (n + 1) / 2

	fmt.Print(sign * mag)
}
