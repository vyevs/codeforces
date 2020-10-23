package main

import "fmt"

// A. Next Round
// time limit per test
// 3 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// "Contestant who earns a score equal to or greater than the k-th place finisher's score will advance to the next round, as long as the contestant earns a positive score..." — an excerpt from contest rules.

// A total of n participants took part in the contest (n ≥ k), and you already know their scores. Calculate how many participants will advance to the next round.
// Input

// The first line of the input contains two integers n and k (1 ≤ k ≤ n ≤ 50) separated by a single space.

// The second line contains n space-separated integers a1, a2, ..., an (0 ≤ ai ≤ 100), where ai is the score earned by the participant who got the i-th place. The given sequence is non-increasing (that is, for all i from 1 to n - 1 the following condition is fulfilled: ai ≥ ai + 1).
// Output

// Output the number of participants who advance to the next round.

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	scores := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&scores[i])
	}

	// k ranges from 1 to n, so subtract 1 because of 0-indexing
	cutoffScore := scores[k-1]

	var nAdvanced int
	// start from 1st place, and keep going until you either hit the end (all contestants advance) or you reach a contestant that will not advance
	// only non-zero scores advance
	for i := 0; i < n && scores[i] != 0 && scores[i] >= cutoffScore; i++ {
		nAdvanced++
	}

	fmt.Print(nAdvanced)
}
