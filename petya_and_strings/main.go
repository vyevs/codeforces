package main

import "fmt"

// A. Petya and Strings
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Little Petya loves presents. His mum bought him two strings of the same size for his birthday. The strings consist of uppercase and lowercase Latin letters. Now Petya wants to compare those two strings lexicographically. The letters' case does not matter, that is an uppercase letter is considered equivalent to the corresponding lowercase letter. Help Petya perform the comparison.
// Input

// Each of the first two lines contains a bought string. The strings' lengths range from 1 to 100 inclusive. It is guaranteed that the strings are of the same length and also consist of uppercase and lowercase Latin letters.
// Output

// If the first string is less than the second one, print "-1". If the second string is less than the first one, print "1". If the strings are equal, print "0". Note that the letters' case is not taken into consideration when the strings are compared.

func main() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)

	var result int
	for i := 0; i < len(str1); i++ {

		// in ascii, the bit 00100000 controls whether an alphabet character is upper case or not (on for lower case, off for upper case)
		// turn this bit off in both string characters for case-insensitive comparison
		c1 := str1[i] & 0xdf
		c2 := str2[i] & 0xdf

		if c1 == c2 {
			continue
		}

		if c1 < c2 {
			result = -1
			break
		} else {
			result = 1
			break
		}
	}

	fmt.Print(result)
}
