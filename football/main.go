package main

import "fmt"

// https://codeforces.com/problemset/problem/96/A
// A. Football
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Petya loves football very much. One day, as he was watching a football match, he was writing the players' current positions on a piece of paper. To simplify the situation he depicted it as a string consisting of zeroes and ones. A zero corresponds to players of one team; a one corresponds to players of another team. If there are at least 7 players of some team standing one after another, then the situation is considered dangerous. For example, the situation 00100110111111101 is dangerous and 11110111011101 is not. You are given the current situation. Determine whether it is dangerous or not.
// Input

// The first input line contains a non-empty string consisting of characters "0" and "1", which represents players. The length of the string does not exceed 100 characters. There's at least one player from each team present on the field.
// Output

// Print "YES" if the situation is dangerous. Otherwise, print "NO".

func main() {
	str := make([]byte, 100)
	_, _ = fmt.Scan(&str)

	var dangerous bool

	streakByte := str[0]
	streak := 1
	for i := 0; i < len(str); i++ {
		b := str[i]

		if b == streakByte {
			streak++
		} else {
			streak = 1
			streakByte = b
		}

		if streak == 7 {
			dangerous = true
			break
		}
	}

	if dangerous {
		_, _ = fmt.Print("YES")
	} else {
		_, _ = fmt.Print("NO")
	}
}
