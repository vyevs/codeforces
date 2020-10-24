package main

import (
	"fmt"
)

// A. Bear and Big Brother
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
// input
// standard input
// output
// standard output

// Bear Limak wants to become the largest of bears, or at least to become larger than his brother Bob.

// Right now, Limak and Bob weigh a and b respectively. It's guaranteed that Limak's weight is smaller than or equal to his brother's weight.

// Limak eats a lot and his weight is tripled after every year, while Bob's weight is doubled after every year.

// After how many full years will Limak become strictly larger (strictly heavier) than Bob?
// Input

// The only line of the input contains two integers a and b (1 ≤ a ≤ b ≤ 10) — the weight of Limak and the weight of Bob respectively.
// Output

// Print one integer, denoting the integer number of years after which Limak will become strictly larger than Bob.

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	// the inequality to solve here is a * (3^y) > b * (2^y) where y is the year
	// solving this inequality using logarithms results in
	// y = (ln(b) - ln(a)) / (ln(3) - ln(2))
	// since this yields a floating point value, we must round it up to the next integer
	// this should theoretically work, but doesn't produce the correct answer in the codeforces environment despite producing correct results on my machine
	// fmt.Print(int(math.Ceil((math.Log(float64(b)) - math.Log(float64(a))) / (math.Log(3) - math.Log(2)))))

	var years int
	for a <= b {
		a *= 3
		b *= 2
		years++
	}
	fmt.Print(years)
}
