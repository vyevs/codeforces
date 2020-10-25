package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/344/A
// A. Magnets
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Mad scientist Mike entertains himself by arranging rows of dominoes. He doesn't need dominoes, though: he uses rectangular magnets instead. Each magnet has two poles, positive (a "plus") and negative (a "minus"). If two magnets are put together at a close distance, then the like poles will repel each other and the opposite poles will attract each other.

// Mike starts by laying one magnet horizontally on the table. During each following step Mike adds one more magnet horizontally to the right end of the row. Depending on how Mike puts the magnet on the table, it is either attracted to the previous one (forming a group of multiple magnets linked together) or repelled by it (then Mike lays this magnet at some distance to the right from the previous one). We assume that a sole magnet not linked to others forms a group of its own.

// Mike arranged multiple magnets in a row. Determine the number of groups that the magnets formed.
// Input

// The first line of the input contains an integer n (1 ≤ n ≤ 100000) — the number of magnets. Then n lines follow. The i-th line (1 ≤ i ≤ n) contains either characters "01", if Mike put the i-th magnet in the "plus-minus" position, or characters "10", if Mike put the magnet in the "minus-plus" position.
// Output

// On the single line of the output print the number of groups of magnets.

func main() {
	solveTheProblem()
}

func solveTheProblem() {
	// this problem has a test case with 100k magnets, so we need to use bufio to reduce the number of IO ops
	// as the 1 second time limit can be very easily be exceeded when you have to read 100k values
	bufStdin := bufio.NewReader(os.Stdin)

	var n int // number of magnets
	fmt.Fscan(bufStdin, &n)

	magnets := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(bufStdin, &magnets[i])
	}

	fmt.Print(howManyGroups(magnets))
}

func howManyGroups(magnets []int) int {
	// by default there is one group since a positive number of magnets can be arranged into a single connected group
	nGroups := 1

	// we look at all consecutive pairs of magnets, and if they face each other with the same poles
	// then the magnets repel and a new "group" is created
	for i := 0; i < len(magnets)-1; i++ {
		// we look at magnets at positions i and i+1

		curMag := magnets[i]
		nextMag := magnets[i+1]

		// the right pole of magnet i is adjcaent to the left pole of magnet i+1
		// if they are the same, the 2 magnets repel and a new group beginning with nextMag is formed
		if curMag != nextMag {
			nGroups++
		}
	}

	return nGroups
}
