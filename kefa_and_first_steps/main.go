package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/580/A
// A. Kefa and First Steps
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Kefa decided to make some money doing business on the Internet for exactly n days. He knows that on the i-th day (1 ≤ i ≤ n) he makes ai money. Kefa loves progress, that's why he wants to know the length of the maximum non-decreasing subsegment in sequence ai. Let us remind you that the subsegment of the sequence is its continuous fragment. A subsegment of numbers is called non-decreasing if all numbers in it follow in the non-decreasing order.

// Help Kefa cope with this task!
// Input

// The first line contains integer n (1 ≤ n ≤ 105).

// The second line contains n integers a1,  a2,  ...,  an (1 ≤ ai ≤ 109).
// Output

// Print a single integer — the length of the maximum non-decreasing subsegment of sequence a.

func main() {
	bufStdin := bufio.NewReader(os.Stdin)

	var n int
	_, _ = fmt.Fscan(bufStdin, &n)

	as := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(bufStdin, &as[i])
	}

	bestStreak := 1
	curStreak := 1
	prevVal := as[0]

	for i := 1; i < n; i++ {
		v := as[i]

		if v >= prevVal {
			curStreak++
			if curStreak > bestStreak {
				bestStreak = curStreak
			}
		} else {
			curStreak = 1
		}

		prevVal = v
	}

	_, _ = fmt.Print(bestStreak)
}
