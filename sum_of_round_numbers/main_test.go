package main

import (
	"strconv"
	"testing"
)

func TestDecomposeIntoRoundNumbers(t *testing.T) {
	tests := []struct {
		in   int
		want []int
	}{
		{
			in:   7,
			want: []int{7},
		},
		{
			in:   17,
			want: []int{10, 7},
		},
		{
			in:   5009,
			want: []int{5000, 9},
		},
		{
			in:   9876,
			want: []int{9000, 800, 70, 6},
		},
		{
			in:   10000,
			want: []int{10000},
		},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.in), func(t *testing.T) {
			got := decomposeIntoRoundNumbers(test.in)

			if !slicesHaveSameValuesInAnyOrder(got, test.want) {
				t.Errorf("got: %v want: %v", got, test.want)
			}
		})
	}
}

func slicesHaveSameValuesInAnyOrder(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1Cts := make(map[int]int, len(s1))
	for _, v := range s1 {
		s1Cts[v]++
	}
	s2Cts := make(map[int]int, len(s2))
	for _, v := range s2 {
		s2Cts[v]++
	}

	for k, v := range s1Cts {
		if s2Cts[k] != v {
			return false
		}
	}
	for k, v := range s2Cts {
		if s1Cts[k] != v {
			return false
		}
	}
	return true
}
