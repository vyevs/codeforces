package main

import "fmt"

// https://codeforces.com/problemset/problem/59/A
// A. Word
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Vasya is very upset that many people on the Net mix uppercase and lowercase letters in one word. That's why he decided to invent an extension for his favorite browser that would change the letters' register in every word so that it either only consisted of lowercase letters or, vice versa, only of uppercase ones. At that as little as possible letters should be changed in the word. For example, the word HoUse must be replaced with house, and the word ViP — with VIP. If a word contains an equal number of uppercase and lowercase letters, you should replace all the letters with lowercase ones. For example, maTRIx should be replaced by matrix. Your task is to use the given method on one given word.
// Input

// The first line contains a word s — it consists of uppercase and lowercase Latin letters and possesses the length from 1 to 100.
// Output

// Print the corrected word s. If the given word s has strictly more uppercase letters, make the word written in the uppercase register, otherwise - in the lowercase one.

func main() {
	var word []byte
	fmt.Scan(&word)

	isLowerCase := func(c byte) bool {
		// the bit 00100000 in an ascii alphabet character controls whether it is lowercase or uppercase
		// the bit is 1 if the character is lower case, so checking this bit is sufficient to tell if the character is lower case
		return (c & 0x20) != 0
	}

	var nLowerCaseLetters int
	for _, c := range word {
		if isLowerCase(c) {
			nLowerCaseLetters++
		}
	}

	if nLowerCaseLetters >= (len(word)+1)/2 {
		// half or more of the letters are lowercase, so lowercase the word
		for i := range word {
			word[i] |= 0x20
		}
	} else {
		// more than half off the letters are uppercase, so uppercase the word
		for i := range word {
			word[i] &= 0xdf
		}
	}

	fmt.Print(string(word))
}
