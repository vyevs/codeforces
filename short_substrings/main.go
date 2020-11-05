package main

// https://codeforces.com/problemset/problem/1367/A
// A. Short Substrings
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Alice guesses the strings that Bob made for her.

// At first, Bob came up with the secret string a
// consisting of lowercase English letters. The string a has a length of 2 or more characters. Then, from string a he builds a new string b and offers Alice the string b so that she can guess the string a

// .

// Bob builds b
// from a as follows: he writes all the substrings of length 2 of the string a in the order from left to right, and then joins them in the same order into the string b

// .

// For example, if Bob came up with the string a
// ="abac", then all the substrings of length 2 of the string a are: "ab", "ba", "ac". Therefore, the string b

// ="abbaac".

// You are given the string b
// . Help Alice to guess the string a that Bob came up with. It is guaranteed that b

// was built according to the algorithm given above. It can be proved that the answer to the problem is unique.
// Input

// The first line contains a single positive integer t
// (1≤t≤1000) — the number of test cases in the test. Then t

// test cases follow.

// Each test case consists of one line in which the string b
// is written, consisting of lowercase English letters (2≤|b|≤100) — the string Bob came up with, where |b| is the length of the string b. It is guaranteed that b

// was built according to the algorithm given above.
// Output

// Output t
// answers to test cases. Each answer is the secret string a, consisting of lowercase English letters, that Bob came up with.


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
        var b []byte
        _, _ = fmt.Fscan(bufStdin, &b)
        
        a := make([]byte, 0, len(b))
        
        // the transformation a -> b duplicates every character in a except for the first and last characters
        // as seen in the example abcd -> abbccd, so we want to deduplicate the letters in the center
        // of b, while keeping the first and last characters
        
        a = append(a, b[0])
        for i := 1; i < len(b)-1; i+=2 {
            a = append(a, b[i])
        }
        a = append(a, b[len(b)-1])
        
        _, _ = fmt.Fprintf(bufStdout, "%s\n", string(a))
    }
}