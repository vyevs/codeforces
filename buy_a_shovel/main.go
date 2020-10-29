package main

import "fmt"

// https://codeforces.com/problemset/problem/732/A
// A. Buy a Shovel
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Polycarp urgently needs a shovel! He comes to the shop and chooses an appropriate one. The shovel that Policarp chooses is sold for k burles. Assume that there is an unlimited number of such shovels in the shop.

// In his pocket Polycarp has an unlimited number of "10-burle coins" and exactly one coin of r burles (1 ≤ r ≤ 9).

// What is the minimum number of shovels Polycarp has to buy so that he can pay for the purchase without any change? It is obvious that he can pay for 10 shovels without any change (by paying the requied amount of 10-burle coins and not using the coin of r burles). But perhaps he can buy fewer shovels and pay without any change. Note that Polycarp should buy at least one shovel.
// Input

// The single line of input contains two integers k and r (1 ≤ k ≤ 1000, 1 ≤ r ≤ 9) — the price of one shovel and the denomination of the coin in Polycarp's pocket that is different from "10-burle coins".

// Remember that he has an unlimited number of coins in the denomination of 10, that is, Polycarp has enough money to buy any number of shovels.
// Output

// Print the required minimum number of shovels Polycarp has to buy so that he can pay for them without any change.

func main() {
	var k, r int
	_, _ = fmt.Scan(&k, &r)

	// the equality that needs to be solve is x*k mod 10 = r or x*k mod 10 = 0, for lowest x

	nShovles := 1
	for {
		price := nShovles * k

		mod := price % 10
		if mod == 0 || mod == r {
			break
		}
		nShovles++
	}

	fmt.Print(nShovles)
}
