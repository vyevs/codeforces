package main

import "fmt"

// https://codeforces.com/problemset/problem/750/A
// A. New Year and Hurry
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Limak is going to participate in a contest on the last day of the 2016. The contest will start at 20:00 and will last four hours, exactly until midnight. There will be n problems, sorted by difficulty, i.e. problem 1 is the easiest and problem n is the hardest. Limak knows it will take him 5·i minutes to solve the i-th problem.

// Limak's friends organize a New Year's Eve party and Limak wants to be there at midnight or earlier. He needs k minutes to get there from his house, where he will participate in the contest first.

// How many problems can Limak solve if he wants to make it to the party?
// Input

// The only line of the input contains two integers n and k (1 ≤ n ≤ 10, 1 ≤ k ≤ 240) — the number of the problems in the contest and the number of minutes Limak needs to get to the party from his house.
// Output

// Print one integer, denoting the maximum possible number of problems Limak can solve so that he could get to the party at midnight or earlier.

func main() {
	var n, k int
	_, _ = fmt.Scan(&n, &k)

	minutesLeft := 240 - k

	var nProblems int
	for i := 0; i < n; i++ {
		timeForProblem := (i + 1) * 5
		minutesLeft -= timeForProblem
		if minutesLeft < 0 {
			break
		}
		nProblems++
	}

	_, _ = fmt.Print(nProblems)
}
