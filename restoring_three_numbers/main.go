package main

import "fmt"

// https://codeforces.com/problemset/problem/1154/A
// A. Restoring Three Numbers
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Polycarp has guessed three positive integers a
// , b and c. He keeps these numbers in secret, but he writes down four numbers on a board in arbitrary order — their pairwise sums (three numbers) and sum of all three numbers (one number). So, there are four numbers on a board in random order: a+b, a+c, b+c and a+b+c

// .

// You have to guess three numbers a
// , b and c

// using given numbers. Print three guessed integers in any order.

// Pay attention that some given numbers a
// , b and c can be equal (it is also possible that a=b=c

// ).
// Input

// The only line of the input contains four positive integers x1,x2,x3,x4
// (2≤xi≤109) — numbers written on a board in random order. It is guaranteed that the answer exists for the given number x1,x2,x3,x4

// .
// Output

// Print such positive integers a
// , b and c that four numbers written on a board are values a+b, a+c, b+c and a+b+c written in some order. Print a, b and c in any order. If there are several answers, you can print any. It is guaranteed that the answer exists.

func main() {
	var x1, x2, x3, x4 int
	_, _ = fmt.Scan(&x1, &x2, &x3, &x4)

	xMax := intMax(x1, x2, x3, x4)

	for _, x := range []int{x1, x2, x3, x4} {
		if x < xMax {
			_, _ = fmt.Print(xMax-x, " ")
		}
	}
}

func intMax(ints ...int) int {
	max := ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] > max {
			max = ints[i]
		}
	}
	return max
}
