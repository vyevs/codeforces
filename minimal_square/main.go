package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1360/A
// A. Minimal Square
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Find the minimum area of a square land on which you can place two identical rectangular a×b

// houses. The sides of the houses should be parallel to the sides of the desired square land.

// Formally,

//     You are given two identical rectangles with side lengths a

// and b (1≤a,b≤100

//     ) — positive integers (you are given just the sizes, but not their positions).
//     Find the square of the minimum area that contains both given rectangles. Rectangles can be rotated (both or just one), moved, but the sides of the rectangles should be parallel to the sides of the desired square.

// Two rectangles can touch each other (side or corner), but cannot intersect. Rectangles can also touch the sides of the square but must be completely inside it. You can rotate the rectangles. Take a look at the examples for a better understanding.
// The picture shows a square that contains red and green rectangles.
// Input

// The first line contains an integer t
// (1≤t≤10000) —the number of test cases in the input. Then t

// test cases follow.

// Each test case is a line containing two integers a
// , b (1≤a,b≤100

// ) — side lengths of the rectangles.
// Output

// Print t
// answers to the test cases. Each answer must be a single integer — minimal area of square land, that contains two rectangles with dimensions a×b.

func main() {
	bufStdin := bufio.NewReader(os.Stdin)
	bufStdout := bufio.NewWriter(os.Stdout)
	defer bufStdout.Flush()

	var t int
	_, _ = fmt.Fscan(bufStdin, &t)

	for i := 0; i < t; i++ {
		var a, b int

		_, _ = fmt.Fscan(bufStdin, &a, &b)

		shorterSide, longerSide := intMin(a, b), intMax(a, b)

		squareSide := intMax(2*shorterSide, longerSide)
		area := squareSide * squareSide
		_, _ = fmt.Fprintln(bufStdout, area)
	}
}

func intMin(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}
func intMax(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}
