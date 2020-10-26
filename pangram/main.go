package main

import "fmt"

// https://codeforces.com/problemset/problem/520/A
// A. Pangram
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// A word or a sentence in some language is called a pangram if all the characters of the alphabet of this language appear in it at least once. Pangrams are often used to demonstrate fonts in printing or test the output devices.

// You are given a string consisting of lowercase and uppercase Latin letters. Check whether this string is a pangram. We say that the string contains a letter of the Latin alphabet if this letter occurs in the string in uppercase or lowercase.
// Input

// The first line contains a single integer n (1 ≤ n ≤ 100) — the number of characters in the string.

// The second line contains the string. The string consists only of uppercase and lowercase Latin letters.
// Output

// Output "YES", if the string is a pangram and "NO" otherwise.

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	var word []byte
	_, _ = fmt.Scan(&word)

	if containsEachAlphabetLetter(word) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

func containsEachAlphabetLetter(word []byte) bool {
	var haveChar [26]bool
	for _, c := range word {
		lowerCaseC := c | 0x20
		haveChar[lowerCaseC-'a'] = true
	}

	for _, b := range haveChar {
		if !b {
			return false
		}
	}
	return true
}
