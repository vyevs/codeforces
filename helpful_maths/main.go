package main

import (
	"fmt"
	"strings"
)

// A. Helpful Maths
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Xenia the beginner mathematician is a third year student at elementary school. She is now learning the addition operation.

// The teacher has written down the sum of multiple numbers. Pupils should calculate the sum. To make the calculation easier, the sum only contains numbers 1, 2 and 3. Still, that isn't enough for Xenia. She is only beginning to count, so she can calculate a sum only if the summands follow in non-decreasing order. For example, she can't calculate sum 1+3+2+1 but she can calculate sums 1+1+2 and 3+3.

// You've got the sum that was written on the board. Rearrange the summans and print the sum in such a way that Xenia can calculate the sum.
// Input

// The first line contains a non-empty string s — the sum Xenia needs to count. String s contains no spaces. It only contains digits and characters "+". Besides, string s is a correct sum of numbers 1, 2 and 3. String s is at most 100 characters long.
// Output

// Print the new sum that Xenia can count.

func main() {
	var line string
	_, _ = fmt.Scan(&line)

	summands := strings.Split(line, "+")

	// one solution is to just sort summands and then join them back with "+"
	// but in order to avoid the sort, just count the number of occurrences of each digit
	// and then write the respective amount of each digit and join it using "+"
	var n1s, n2s, n3s int
	for _, summand := range summands {
		if summand == "1" {
			n1s++
		} else if summand == "2" {
			n2s++
		} else if summand == "3" {
			n3s++
		}
	}

	for i := 0; i < n1s; i++ {
		summands[i] = "1"
	}
	for i := 0; i < n2s; i++ {
		summands[n1s+i] = "2"
	}
	for i := 0; i < n3s; i++ {
		summands[n1s+n2s+i] = "3"
	}

	fmt.Print(strings.Join(summands, "+"))
}
