package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/703/A
// Rules of the game are very simple: at first number of rounds n is defined. In every round each of the players throws a cubical dice with distinct numbers from 1 to 6 written on its faces. Player, whose value after throwing the dice is greater, wins the round. In case if player dice values are equal, no one of them is a winner.

// In average, player, who won most of the rounds, is the winner of the game. In case if two players won the same number of rounds, the result of the game is draw.

// Mishka is still very little and can't count wins and losses, so she asked you to watch their game and determine its result. Please help her!
// Input

// The first line of the input contains single integer n n (1 ≤ n ≤ 100) — the number of game rounds.

// The next n lines contains rounds description. i-th of them contains pair of integers mi and ci (1 ≤ mi,  ci ≤ 6) — values on dice upper face after Mishka's and Chris' throws in i-th round respectively.
// Output

// If Mishka is the winner of the game, print "Mishka" (without quotes) in the only line.

// If Chris is the winner of the game, print "Chris" (without quotes) in the only line.

// If the result of the game is draw, print "Friendship is magic!^^" (without quotes) in the only line.

func main() {
	bufStdin := bufio.NewReader(os.Stdin)

	var n int
	_, _ = fmt.Fscan(bufStdin, &n)

	var mWins, cWins int
	for i := 0; i < n; i++ {
		var mi, ci int
		_, _ = fmt.Fscan(bufStdin, &mi, &ci)

		if mi > ci {
			mWins++
		} else if ci > mi {
			cWins++
		}
	}

	if mWins > cWins {
		_, _ = fmt.Print("Mishka")
	} else if cWins > mWins {
		_, _ = fmt.Print("Chris")
	} else {
		_, _ = fmt.Print("Friendship is magic!^^")
	}
}
