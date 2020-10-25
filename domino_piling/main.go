package main

import "fmt"

// https://codeforces.com/problemset/problem/50/A
// A. Domino piling
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// You are given a rectangular board of M × N squares. Also you are given an unlimited number of standard domino pieces of 2 × 1 squares. You are allowed to rotate the pieces. You are asked to place as many dominoes as possible on the board so as to meet the following conditions:

// 1. Each domino completely covers two squares.

// 2. No two dominoes overlap.

// 3. Each domino lies entirely inside the board. It is allowed to touch the edges of the board.

// Find the maximum number of dominoes, which can be placed under these restrictions.
// Input

// In a single line you are given two integers M and N — board sizes in squares (1 ≤ M ≤ N ≤ 16).
// Output

// Output one number — the maximal number of dominoes, which can be placed.

func main() {
	var m, n int
	fmt.Scan(&m, &n)

	// the solution to this is to greedily place dominoes, whether you place them vertically or horizontally doesn't matter
	// if we say the grid is mxn and m is the horizontal axis and n is the vertical axis
	// (m/2) * n first attempts to place dominoes horizontally
	//  m/2 is how many dominoes we can place horizontally, and n is the number of rows of these horizontal placements we can do
	//
	// ((n / 2) * (m % 2)) attempts to place dominoes vertically in the n direction
	// n/2 is how many dominoes we can place vertically, and m%2 is the number of columns we have left after placing the horizontal dominoes
	nDominoes := ((m / 2) * n) + ((n / 2) * (m % 2))

	fmt.Print(nDominoes)
}
