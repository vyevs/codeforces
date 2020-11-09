package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/158/B
// B. Taxi
// time limit per test
// 3 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// After the lessons n groups of schoolchildren went outside and decided to visit Polycarpus to celebrate his birthday. We know that the i-th group consists of si friends (1 ≤ si ≤ 4), and they want to go to Polycarpus together. They decided to get there by taxi. Each car can carry at most four passengers. What minimum number of cars will the children need if all members of each group should ride in the same taxi (but one taxi can take more than one group)?
// Input

// The first line contains integer n (1 ≤ n ≤ 105) — the number of groups of schoolchildren. The second line contains a sequence of integers s1, s2, ..., sn (1 ≤ si ≤ 4). The integers are separated by a space, si is the number of children in the i-th group.
// Output

// Print the single number — the minimum number of taxis necessary to drive all children to Polycarpus.

func main() {
	bufStdin := bufio.NewReader(os.Stdin)

	var n int
	_, _ = fmt.Fscan(bufStdin, &n)

	ss := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(bufStdin, &ss[i])
	}

	low, high := 0, len(ss)-1

	for low <= high {

	}

	_, _ = fmt.Print(nCars)
}
