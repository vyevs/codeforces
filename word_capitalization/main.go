package main

import "fmt"

// https://codeforces.com/problemset/problem/281/A
// A. Word Capitalization
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Capitalization is writing a word with its first letter as a capital letter. Your task is to capitalize the given word.

// Note, that during capitalization all the letters except the first one remains unchanged.
// Input

// A single line contains a non-empty word. This word consists of lowercase and uppercase English letters. The length of the word will not exceed 103.
// Output

// Output the given word after capitalization.

func main() {
	var word []byte
	fmt.Scan(&word)

	// the bit 00100000 controls whether an ascii alphabet character is uppercase, on for lower case, off for upper case
	// ANDing against 0xdf turns the bit off if it was on, and has no effect otherwise, capitalizing the first character
	word[0] &= 0xdf

	fmt.Print(string(word))
}
