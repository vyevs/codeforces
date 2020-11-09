package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// https://codeforces.com/problemset/problem/160/A
// A. Twins
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Imagine that you have a twin brother or sister. Having another person that looks exactly like you seems very unusual. It's hard to say if having something of an alter ego is good or bad. And if you do have a twin, then you very well know what it's like.

// Now let's imagine a typical morning in your family. You haven't woken up yet, and Mom is already going to work. She has been so hasty that she has nearly forgotten to leave the two of her darling children some money to buy lunches in the school cafeteria. She fished in the purse and found some number of coins, or to be exact, n coins of arbitrary values a1, a2, ..., an. But as Mom was running out of time, she didn't split the coins for you two. So she scribbled a note asking you to split the money equally.

// As you woke up, you found Mom's coins and read her note. "But why split the money equally?" — you thought. After all, your twin is sleeping and he won't know anything. So you decided to act like that: pick for yourself some subset of coins so that the sum of values of your coins is strictly larger than the sum of values of the remaining coins that your twin will have. However, you correctly thought that if you take too many coins, the twin will suspect the deception. So, you've decided to stick to the following strategy to avoid suspicions: you take the minimum number of coins, whose sum of values is strictly more than the sum of values of the remaining coins. On this basis, determine what minimum number of coins you need to take to divide them in the described manner.
// Input

// The first line contains integer n (1 ≤ n ≤ 100) — the number of coins. The second line contains a sequence of n integers a1, a2, ..., an (1 ≤ ai ≤ 100) — the coins' values. All numbers are separated with spaces.
// Output

// In the single line print the single number — the minimum needed number of coins.

func main() {
	bufStdin := bufio.NewReader(os.Stdin)

	var n int
	_, _ = fmt.Fscan(bufStdin, &n)

	coins := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(bufStdin, &coins[i])
	}

	sort.Ints(coins)

	var leftSum int
	for i := 0; i < len(coins)-1; i++ {
		leftSum += coins[i]
	}

	rightSum := coins[len(coins)-1]

	nCoins := 1

	for rightSum <= leftSum {
		exchange := coins[len(coins)-nCoins-1]

		leftSum -= exchange
		rightSum += exchange

		nCoins++
	}

	_, _ = fmt.Print(nCoins)

}
