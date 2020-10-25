package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://codeforces.com/problemset/problem/705/A
// A. Hulk
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Dr. Bruce Banner hates his enemies (like others don't). As we all know, he can barely talk when he turns into the incredible Hulk. That's why he asked you to help him to express his feelings.

// Hulk likes the Inception so much, and like that his feelings are complicated. They have n layers. The first layer is hate, second one is love, third one is hate and so on...

// For example if n = 1, then his feeling is "I hate it" or if n = 2 it's "I hate that I love it", and if n = 3 it's "I hate that I love that I hate it" and so on.

// Please help Dr. Banner.
// Input

// The only line of the input contains a single integer n (1 ≤ n ≤ 100) — the number of layers of love and hate.
// Output

// Print Dr.Banner's feeling in one line.

func main() {
	solveTheProblem()
}

func solveTheProblem() {
	bufStdin := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(bufStdin, &n)

	var feeling strings.Builder
	feeling.Grow(n * 16) // this is always enough bytes to write the entire string so we never have to do allocations while writing the result

	// 1 <= n <= 100 so we always begin with an "I hate"
	feeling.WriteString("I hate")

	for i := 1; i < n; i++ {
		feeling.WriteString(" that")
		if i%2 == 0 {
			feeling.WriteString(" I hate")
		} else {
			feeling.WriteString(" I love")
		}
	}

	// we always end with "it"
	feeling.WriteString(" it")

	fmt.Print(feeling.String())
}
