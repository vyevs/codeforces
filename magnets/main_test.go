package main

import "testing"

func BenchmarkHowManyGroups(b *testing.B) {
	magnets := make100kMagnets()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		nGroups := howManyGroups(magnets)
		if nGroups != 100_000 {
			b.Fatalf("got wrong answer %d, want %d", nGroups, 100_000)
		}
	}
}

func make100kMagnets() []int {
	magnets := make([]int, 100_000)
	for i := 0; i < 100_000; i++ {
		if i%2 == 0 {
			magnets[i] = 10
		} else {
			magnets[i] = 1
		}
	}
	return magnets
}
