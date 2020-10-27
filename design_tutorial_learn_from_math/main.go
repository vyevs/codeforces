package main

import "fmt"

// https://codeforces.com/problemset/problem/472/A
// A. Design Tutorial: Learn from Math
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// One way to create a task is to learn from math. You can generate some random math statement or modify some theorems to get something new and build a new task from that.

// For example, there is a statement called the "Goldbach's conjecture". It says: "each even number no less than four can be expressed as the sum of two primes". Let's modify it. How about a statement like that: "each integer no less than 12 can be expressed as the sum of two composite numbers." Not like the Goldbach's conjecture, I can prove this theorem.

// You are given an integer n no less than 12, express it as a sum of two composite numbers.
// Input

// The only line contains an integer n (12 ≤ n ≤ 106).
// Output

// Output two composite integers x and y (1 < x, y < n) such that x + y = n. If there are multiple solutions, you can output any of them.

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	// there is a neat trick here that allows us to determine the composite pair in a trivial way
	// if n is odd, 9 and n-9 are a composite pair
	// n-9 is one of the numbers in the sequence 4, 6, 8, 10, 12, etc... so it is composite
	//
	// if n is even, 8 and n-8 are a composite pair
	// n-8 is one of the numbers in the sequence 4, 6, 8, 10, 12, etc... so it is composite

	if n%2 == 0 {
		fmt.Print(8, n-8)
	} else {
		fmt.Print(9, n-9)
	}
}
