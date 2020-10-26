package main

import "fmt"

// https://codeforces.com/problemset/problem/996/A
// A. Hit the Lottery
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Allen has a LOT of money. He has n
// dollars in the bank. For security reasons, he wants to withdraw it in cash (we will not disclose the reasons here). The denominations for dollar bills are 1, 5, 10, 20, 100

// . What is the minimum number of bills Allen could receive after withdrawing his entire balance?
// Input

// The first and only line of input contains a single integer n
// (1≤n≤109

// ).
// Output

// Output the minimum number of bills that Allen could receive.

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	var nBills int

	nBills += n / 100
	n %= 100

	nBills += n / 20
	n %= 20

	nBills += n / 10
	n %= 10

	nBills += n / 5
	n %= 5

	nBills += n

	fmt.Print(nBills)
}
