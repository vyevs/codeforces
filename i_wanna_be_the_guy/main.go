package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/469/A
// A. I Wanna Be the Guy
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// There is a game called "I Wanna Be the Guy", consisting of n levels. Little X and his friend Little Y are addicted to the game. Each of them wants to pass the whole game.

// Little X can pass only p levels of the game. And Little Y can pass only q levels of the game. You are given the indices of levels Little X can pass and the indices of levels Little Y can pass. Will Little X and Little Y pass the whole game, if they cooperate each other?
// Input

// The first line contains a single integer n (1 ≤  n ≤ 100).

// The next line contains an integer p (0 ≤ p ≤ n) at first, then follows p distinct integers a1, a2, ..., ap (1 ≤ ai ≤ n). These integers denote the indices of levels Little X can pass. The next line contains the levels Little Y can pass in the same format. It's assumed that levels are numbered from 1 to n.
// Output

// If they can pass all the levels, print "I become the guy.". If it's impossible, print "Oh, my keyboard!" (without the quotes).

func main() {
	solveTheProblem()
}

func solveTheProblem() {
	bufStdin := bufio.NewReader(os.Stdin)

	var n int // level we need to pass
	_, _ = fmt.Fscan(bufStdin, &n)

	var p int // number of levels Little X can pass
	_, _ = fmt.Fscan(bufStdin, &p)
	xLevels := make([]int, p)
	for i := 0; i < p; i++ {
		_, _ = fmt.Fscan(bufStdin, &xLevels[i])
	}

	var q int // number of levels Little Y can pass
	_, _ = fmt.Fscan(bufStdin, &q)
	yLevels := make([]int, q)
	for i := 0; i < q; i++ {
		_, _ = fmt.Fscan(bufStdin, &yLevels[i])
	}

	if canPassLevels(xLevels, yLevels, n) {
		fmt.Print("I become the guy.")
	} else {
		fmt.Print("Oh, my keyboard!")
	}
}

// this is the simplest solution, use map as a set to record which levels we can pass, and then check each level
func canPassLevels(xLevels, yLevels []int, n int) bool {
	levelsWeCanPass := make(map[int]struct{}, len(xLevels)+len(yLevels))

	for _, v := range xLevels {
		levelsWeCanPass[v] = struct{}{}
	}
	for _, v := range yLevels {
		levelsWeCanPass[v] = struct{}{}
	}

	for i := 1; i <= n; i++ {
		if _, canPass := levelsWeCanPass[i]; !canPass {
			return false
		}
	}

	return true
}
