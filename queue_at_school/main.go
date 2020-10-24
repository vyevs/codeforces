package main

import "fmt"

// B. Queue at the School
// time limit per test
// 2 seconds
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// During the break the schoolchildren, boys and girls, formed a queue of n people in the canteen. Initially the children stood in the order they entered the canteen. However, after a while the boys started feeling awkward for standing in front of the girls in the queue and they started letting the girls move forward each second.

// Let's describe the process more precisely. Let's say that the positions in the queue are sequentially numbered by integers from 1 to n, at that the person in the position number 1 is served first. Then, if at time x a boy stands on the i-th position and a girl stands on the (i + 1)-th position, then at time x + 1 the i-th position will have a girl and the (i + 1)-th position will have a boy. The time is given in seconds.

// You've got the initial position of the children, at the initial moment of time. Determine the way the queue is going to look after t seconds.
// Input

// The first line contains two integers n and t (1 ≤ n, t ≤ 50), which represent the number of children in the queue and the time after which the queue will transform into the arrangement you need to find.

// The next line contains string s, which represents the schoolchildren's initial arrangement. If the i-th position in the queue contains a boy, then the i-th character of string s equals "B", otherwise the i-th character equals "G".
// Output

// Print string a, which describes the arrangement after t seconds. If the i-th position has a boy after the needed time, then the i-th character a must equal "B", otherwise it must equal "G"

func main() {
	var n, t int
	var queue string
	fmt.Scan(&n, &t, &queue)

	queueBytes := []byte(queue)

	for i := 0; i < t; i++ {
		for i := 0; i < len(queueBytes)-1; {
			if queueBytes[i] == 'B' && queueBytes[i+1] == 'G' {
				queueBytes[i] = 'G'
				queueBytes[i+1] = 'B'
				i += 2
				// if we did a swap, we can jump past both of the swapped places
				// since we don't want to keep doing swaps in a situation like BGG, the end result should be GBG, not GGB
			} else {
				i++
			}
		}
	}

	fmt.Print(string(queueBytes))
}
