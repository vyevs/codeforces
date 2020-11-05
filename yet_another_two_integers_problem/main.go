package main

// https://codeforces.com/problemset/problem/1409/A
// A. Yet Another Two Integers Problem
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// You are given two integers a
// and b

// .

// In one move, you can choose some integer k
// from 1 to 10 and add it to a or subtract it from a. In other words, you choose an integer k∈[1;10] and perform a:=a+k or a:=a−k. You may use different values of k

// in different moves.

// Your task is to find the minimum number of moves required to obtain b
// from a

// .

// You have to answer t

// independent test cases.
// Input

// The first line of the input contains one integer t
// (1≤t≤2⋅104) — the number of test cases. Then t

// test cases follow.

// The only line of the test case contains two integers a
// and b (1≤a,b≤109

// ).
// Output

// For each test case, print the answer: the minimum number of moves required to obtain b
// from a.

import (
    "bufio"
    "fmt"
    "os"
)

var (
    bufStdin  = bufio.NewReader(os.Stdin)
    bufStdout = bufio.NewWriter(os.Stdout)
)

func main() {
    defer bufStdout.Flush()
    
    var t int
    _, _ = fmt.Fscan(bufStdin, &t)
    
    for i := 0; i < t; i++ {
        var a, b int
        _, _ = fmt.Fscan(bufStdin, &a, &b)
        
        
        // the minimum number of steps from a to b is the absolute difference divided by 10 (using integer division) plus 1 if there is a remainder
        // we want to take the largest possible steps (of size 10) until we need to take a step of less than 10
        // e.g. 78 -> 112 we go 78 -> 88 -> 98 -> 108 -> 114, 4 steps
        // 112 - 78 = 34, 34/10 = 3, since there is a remainder of this division (4), that is one more step we must take
        
        absDiff := intAbs(a - b)
        
        steps := absDiff / 10 // size 10 steps
        if absDiff % 10 != 0 {
            steps++ // final step if the size 10 steps did not suffice
        }
        
        _, _ = fmt.Fprintln(bufStdout, steps)
    }
}

func intAbs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}