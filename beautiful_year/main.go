package main

import "fmt"

// https://codeforces.com/problemset/problem/271/A
// A. Beautiful Year
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// It seems like the year of 2013 came only yesterday. Do you know a curious fact? The year of 2013 is the first year after the old 1987 with only distinct digits.

// Now you are suggested to solve the following problem: given a year number, find the minimum year number which is strictly larger than the given one and has only distinct digits.
// Input

// The single line contains integer y (1000 ≤ y ≤ 9000) — the year number.
// Output

// Print a single integer — the minimum year number that is strictly larger than y and all it's digits are distinct. It is guaranteed that the answer exists.

func main() {
	var y int
	fmt.Scan(&y)

	fmt.Print(nextLuckyYear(y))
}

func nextLuckyYear(y int) int {
	// pretty straightforward, increment the year and check if it's a beautiful year, if it is, then that is the answer
	for candidateY := y + 1; ; candidateY++ {
		if isBeautifulYear(candidateY) {
			return candidateY
		}
	}
}

func isBeautifulYear(y int) bool {
	var seenDigit [10]bool

	for ; y > 0; y /= 10 {
		digit := y % 10
		if seenDigit[digit] {
			return false
		}
		seenDigit[digit] = true
	}
	return true
}
