package main

import "fmt"

// https://codeforces.com/problemset/problem/141/A
// A. Amusing Joke
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// So, the New Year holidays are over. Santa Claus and his colleagues can take a rest and have guests at last. When two "New Year and Christmas Men" meet, thear assistants cut out of cardboard the letters from the guest's name and the host's name in honor of this event. Then the hung the letters above the main entrance. One night, when everyone went to bed, someone took all the letters of our characters' names. Then he may have shuffled the letters and put them in one pile in front of the door.

// The next morning it was impossible to find the culprit who had made the disorder. But everybody wondered whether it is possible to restore the names of the host and his guests from the letters lying at the door? That is, we need to verify that there are no extra letters, and that nobody will need to cut more letters.

// Help the "New Year and Christmas Men" and their friends to cope with this problem. You are given both inscriptions that hung over the front door the previous night, and a pile of letters that were found at the front door next morning.
// Input

// The input file consists of three lines: the first line contains the guest's name, the second line contains the name of the residence host and the third line contains letters in a pile that were found at the door in the morning. All lines are not empty and contain only uppercase Latin letters. The length of each line does not exceed 100.
// Output

// Print "YES" without the quotes, if the letters in the pile could be permuted to make the names of the "New Year and Christmas Men". Otherwise, print "NO" without the quotes.

func main() {
	var name1, name2, letters []byte
	_, _ = fmt.Scan(&name1, &name2, &letters)

	var nameLetterCounts [26]int
	for _, c := range name1 {
		nameLetterCounts[c-'A']++
	}
	for _, c := range name2 {
		nameLetterCounts[c-'A']++
	}

	var lettersLetterCounts [26]int
	for _, c := range letters {
		lettersLetterCounts[c-'A']++
	}

	countsAreEqual := func(c1 []int, c2 []int) bool {
		for i, v := range c1 {
			if v != c2[i] {
				return false
			}
		}
		return true
	}

	if countsAreEqual(nameLetterCounts[:], lettersLetterCounts[:]) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
