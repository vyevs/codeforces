package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/268/A
// A. Games
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Manao works on a sports TV. He's spent much time watching the football games of some country. After a while he began to notice different patterns. For example, each team has two sets of uniforms: home uniform and guest uniform. When a team plays a game at home, the players put on the home uniform. When a team plays as a guest on somebody else's stadium, the players put on the guest uniform. The only exception to that rule is: when the home uniform color of the host team matches the guests' uniform, the host team puts on its guest uniform as well. For each team the color of the home and guest uniform is different.

// There are n teams taking part in the national championship. The championship consists of n·(n - 1) games: each team invites each other team to its stadium. At this point Manao wondered: how many times during the championship is a host team going to put on the guest uniform? Note that the order of the games does not affect this number.

// You know the colors of the home and guest uniform for each team. For simplicity, the colors are numbered by integers in such a way that no two distinct colors have the same number. Help Manao find the answer to his question.
// Input

// The first line contains an integer n (2 ≤ n ≤ 30). Each of the following n lines contains a pair of distinct space-separated integers hi, ai (1 ≤ hi, ai ≤ 100) — the colors of the i-th team's home and guest uniforms, respectively.
// Output

// In a single line print the number of games where the host team is going to play in the guest uniform.

func main() {
	bufStdin := bufio.NewReader(os.Stdin)

	var n int
	_, _ = fmt.Fscan(bufStdin, &n)

	homeUniforms := make([]int, n)
	awayUniforms := make([]int, n)

	for i := 0; i < n; i++ {
		var hi, ai int
		_, _ = fmt.Fscan(bufStdin, &hi, &ai)
		homeUniforms[i] = hi
		awayUniforms[i] = ai
	}

	// due to the small range of number of teams present, the double loop solution works fine
	var answer int
	for i, homeUniform := range homeUniforms {
		for j, awayUniform := range awayUniforms {
			if i != j && homeUniform == awayUniform {
				answer++
			}
		}
	}

	fmt.Print(answer)
}
