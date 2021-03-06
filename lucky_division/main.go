package main

import "fmt"

// https://codeforces.com/problemset/problem/122/A
// A. Lucky Division
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Petya loves lucky numbers. Everybody knows that lucky numbers are positive integers whose decimal representation contains only the lucky digits 4 and 7. For example, numbers 47, 744, 4 are lucky and 5, 17, 467 are not.

// Petya calls a number almost lucky if it could be evenly divided by some lucky number. Help him find out if the given number n is almost lucky.
// Input

// The single line contains an integer n (1 ≤ n ≤ 1000) — the number that needs to be checked.
// Output

// In the only line print "YES" (without the quotes), if number n is almost lucky. Otherwise, print "NO" (without the quotes).

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	if isAlmostLucky(n) {
		_, _ = fmt.Print("YES")
	} else {
		_, _ = fmt.Print("NO")
	}
}

func isAlmostLucky(n int) bool {
	if isLucky(n) {
		return true
	}

	halfN := n / 2

	for i := 4; i <= halfN; i++ {
		if !isLucky(i) {
			continue
		}

		if n%i == 0 {
			return true
		}
	}

	return false
}

func isLucky(n int) bool {
	for n > 0 {
		d := n % 10

		if d != 4 && d != 7 {
			return false
		}

		n /= 10
	}

	return true
}
