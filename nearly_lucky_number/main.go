package main

import "fmt"

// https://codeforces.com/problemset/problem/110/A
// A. Nearly Lucky Number
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Petya loves lucky numbers. We all know that lucky numbers are the positive integers whose decimal representations contain only the lucky digits 4 and 7. For example, numbers 47, 744, 4 are lucky and 5, 17, 467 are not.

// Unfortunately, not all numbers are lucky. Petya calls a number nearly lucky if the number of lucky digits in it is a lucky number. He wonders whether number n is a nearly lucky number.
// Input

// The only line contains an integer n (1 ≤ n ≤ 1018).

// Please do not use the %lld specificator to read or write 64-bit numbers in С++. It is preferred to use the cin, cout streams or the %I64d specificator.
// Output

// Print on the single line "YES" if n is a nearly lucky number. Otherwise, print "NO" (without the quotes).

func main() {
	var n uint64
	fmt.Scan(&n)

	var nLuckyDigits int
	for ; n > 0; n /= 10 {
		digit := n % 10
		if digit == 4 || digit == 7 {
			nLuckyDigits++
		}
	}

	if nLuckyDigits == 0 {
		fmt.Print("NO")
		return
	}

	for ; nLuckyDigits > 0; nLuckyDigits /= 10 {
		digit := nLuckyDigits % 10
		if digit != 4 && digit != 7 {
			fmt.Print("NO")
			return
		}
	}

	fmt.Print("YES")
}
