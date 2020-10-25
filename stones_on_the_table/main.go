package main

import "fmt"

// https://codeforces.com/problemset/problem/266/A
// A. Stones on the Table
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// There are n stones on the table in a row, each of them can be red, green or blue. Count the minimum number of stones to take from the table so that any two neighboring stones had different colors. Stones in a row are considered neighboring if there are no other stones between them.
// Input

// The first line contains integer n (1 ≤ n ≤ 50) — the number of stones on the table.

// The next line contains string s, which represents the colors of the stones. We'll consider the stones in the row numbered from 1 to n from left to right. Then the i-th character s equals "R", if the i-th stone is red, "G", if it's green and "B", if it's blue.
// Output

// Print a single integer — the answer to the problem.

func main() {
	{
		// we don't actually need the number of stones, the length of the string we read is enough
		var nStones int
		fmt.Scan(&nStones)
	}

	var stones string
	fmt.Scan(&stones)

	if len(stones) == 1 {
		fmt.Print(0)
		return
	}

	// if we remove a stone at position p, stones at positions p-1 and p+1 are adjacent, and may be the same color!
	// this requires that we keep removing stones until the 2 adjacent stones we are looking at are no longer the same color
	// first and next are the 2 positions we are examining during any one iteration
	first := 0
	next := 1
	var nRemoves int
	for next < len(stones) {
		if stones[first] == stones[next] {
			// if the 2 stones we are looking at are adjacent, "remove" the 2nd of the two by advancing next
			nRemoves++
			next++
		} else {
			// if the 2 stones we are looking at are not adjacent, we can continue onto the next pair
			// since we may have "removed" certain stones by advancing next, we want to move onto the pair next, next+1
			first = next
			next = first + 1
		}
	}

	// sample, f is position of first, n is position of next, the 3 middle stones are removed
	// fn    f n   f  n  f   n f    n
	// RRRRR RXRRR RXXRR RXXXR RXXXR
	//
	// 1 is removed
	// fn          fn          fn          fn          fn          fn          fn         f n          fn
	// BRBGBGRRBR BRBGBGRRBR BRBGBGRRBR BRBGBGRRBR BRBGBGRRBR BRBGBGRRBR BRBGBGRRBR BRBGBGRXBR BRBGBGRXBR

	fmt.Print(nRemoves)
}
