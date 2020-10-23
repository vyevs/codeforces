package main

import "fmt"

// A. Soldier and Bananas
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// A soldier wants to buy w bananas in the shop. He has to pay k dollars for the first banana, 2k dollars for the second one and so on (in other words, he has to pay i·k dollars for the i-th banana).

// He has n dollars. How many dollars does he have to borrow from his friend soldier to buy w bananas?
// Input

// The first line contains three positive integers k, n, w (1  ≤  k, w  ≤  1000, 0 ≤ n ≤ 109), the cost of the first banana, initial number of dollars the soldier has and number of bananas he wants.
// Output

// Output one integer — the amount of dollars that the soldier must borrow from his friend. If he doesn't have to borrow money, output 0.

func main() {
	// k is the cost of the 1st banana
	// n is the number of dollars we start with
	// w is the number of bananas would like to purchase
	var k, n, w int
	fmt.Scan(&k, &n, &w)

	// this is the total amount of money we need to buy the w bananas
	// it uses Gauss' formula ( (w*(w-1)) / 2) ) to calculate the sum of the numbers 1 through w
	// because the costs of the bananas is k for the 1st one, 2k for the 2nd one, 3k, 4k, etc.
	// (1k + 2k + 3k + 4k + ...) = (1 + 2 + 3 + 4 + ...) * k
	totalCost := ((w * (w + 1)) / 2) * k

	if n >= totalCost {
		fmt.Print(0)
	} else {
		fmt.Print(totalCost - n)
	}
}
