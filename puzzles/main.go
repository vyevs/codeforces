package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// https://codeforces.com/problemset/problem/337/A
// A. Puzzles
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// The end of the school year is near and Ms. Manana, the teacher, will soon have to say goodbye to a yet another class. She decided to prepare a goodbye present for her n students and give each of them a jigsaw puzzle (which, as wikipedia states, is a tiling puzzle that requires the assembly of numerous small, often oddly shaped, interlocking and tessellating pieces).

// The shop assistant told the teacher that there are m puzzles in the shop, but they might differ in difficulty and size. Specifically, the first jigsaw puzzle consists of f1 pieces, the second one consists of f2 pieces and so on.

// Ms. Manana doesn't want to upset the children, so she decided that the difference between the numbers of pieces in her presents must be as small as possible. Let A be the number of pieces in the largest puzzle that the teacher buys and B be the number of pieces in the smallest such puzzle. She wants to choose such n puzzles that A - B is minimum possible. Help the teacher and find the least possible value of A - B.
// Input

// The first line contains space-separated integers n and m (2 ≤ n ≤ m ≤ 50). The second line contains m space-separated integers f1, f2, ..., fm (4 ≤ fi ≤ 1000) — the quantities of pieces in the puzzles sold in the shop.
// Output

// Print a single integer — the least possible difference the teacher can obtain.

func main() {
	bufStdin := bufio.NewReader(os.Stdin)

	var n, m int
	_, _ = fmt.Fscan(bufStdin, &n, &m)

	fs := make([]int, m)
	for i := 0; i < m; i++ {
		_, _ = fmt.Fscan(bufStdin, &fs[i])
	}

	sort.Ints(fs)

	min := fs[n-1] - fs[0]
	for i := 1; i <= m-n; i++ {
		s := fs[n-1+i] - fs[i]
		if s < min {
			min = s
		}
	}

	_, _ = fmt.Print(min)
}
