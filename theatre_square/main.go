package main

import "fmt"

// https://codeforces.com/problemset/problem/1/A
// A. Theatre Square
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Theatre Square in the capital city of Berland has a rectangular shape with the size n × m meters. On the occasion of the city's anniversary, a decision was taken to pave the Square with square granite flagstones. Each flagstone is of the size a × a.

// What is the least number of flagstones needed to pave the Square? It's allowed to cover the surface larger than the Theatre Square, but the Square has to be covered. It's not allowed to break the flagstones. The sides of flagstones should be parallel to the sides of the Square.
// Input

// The input contains three positive integer numbers in the first line: n,  m and a (1 ≤  n, m, a ≤ 109).
// Output

// Write the needed number of flagstones.

func main() {
	var n, m, a uint64
	_, _ = fmt.Scan(&n, &m, &a)

	horizontalSquares := n / a
	if n%a != 0 {
		horizontalSquares++
	}

	verticalSquares := m / a
	if m%a != 0 {
		verticalSquares++
	}

	_, _ = fmt.Print(horizontalSquares * verticalSquares)
}
