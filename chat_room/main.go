package main

import "fmt"

// https://codeforces.com/problemset/problem/58/A
// A. Chat room
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Vasya has recently learned to type and log on to the Internet. He immediately entered a chat room and decided to say hello to everybody. Vasya typed the word s. It is considered that Vasya managed to say hello if several letters can be deleted from the typed word so that it resulted in the word "hello". For example, if Vasya types the word "ahhellllloou", it will be considered that he said hello, and if he types "hlelo", it will be considered that Vasya got misunderstood and he didn't manage to say hello. Determine whether Vasya managed to say hello by the given word s.
// Input

// The first and only line contains the word s, which Vasya typed. This word consisits of small Latin letters, its length is no less that 1 and no more than 100 letters.
// Output

// If Vasya managed to say hello, print "YES", otherwise print "NO".

func main() {
	s := make([]byte, 100)
	_, _ = fmt.Scan(&s)

	var saidHello bool

	targetChars := []byte("hello")
	var i int
	for _, b := range s {
		if b == targetChars[i] {
			i++
		}

		if i == 5 {
			saidHello = true
			break
		}
	}

	if saidHello {
		_, _ = fmt.Print("YES")
	} else {
		_, _ = fmt.Print("NO")
	}
}
