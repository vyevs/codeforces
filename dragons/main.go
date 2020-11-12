package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// https://codeforces.com/problemset/problem/230/A
// A. Dragons
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Kirito is stuck on a level of the MMORPG he is playing now. To move on in the game, he's got to defeat all n dragons that live on this level. Kirito and the dragons have strength, which is represented by an integer. In the duel between two opponents the duel's outcome is determined by their strength. Initially, Kirito's strength equals s.

// If Kirito starts duelling with the i-th (1 ≤ i ≤ n) dragon and Kirito's strength is not greater than the dragon's strength xi, then Kirito loses the duel and dies. But if Kirito's strength is greater than the dragon's strength, then he defeats the dragon and gets a bonus strength increase by yi.

// Kirito can fight the dragons in any order. Determine whether he can move on to the next level of the game, that is, defeat all dragons without a single loss.
// Input

// The first line contains two space-separated integers s and n (1 ≤ s ≤ 104, 1 ≤ n ≤ 103). Then n lines follow: the i-th line contains space-separated integers xi and yi (1 ≤ xi ≤ 104, 0 ≤ yi ≤ 104) — the i-th dragon's strength and the bonus for defeating it.
// Output

// On a single line print "YES" (without the quotes), if Kirito can move on to the next level and print "NO" (without the quotes), if he can't.

func main() {
	bufStdin := bufio.NewReader(os.Stdin)

	var s, n int
	_, _ = fmt.Fscan(bufStdin, &s, &n)

	dragons := make([]dragon, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(bufStdin, &dragons[i].str, &dragons[i].bonus)
	}

	sort.Slice(dragons, func(i, j int) bool {
		return dragons[i].str < dragons[j].str
	})

	if canDefeatAll(s, dragons) {
		_, _ = fmt.Print("YES")
	} else {
		_, _ = fmt.Print("NO")
	}
}

type dragon struct {
	str, bonus int
}

func canDefeatAll(s int, ds []dragon) bool {
	for _, d := range ds {
		if s <= d.str {
			return false
		}
		s += d.bonus
	}
	return true
}
