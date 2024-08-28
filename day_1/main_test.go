package main

import "testing"

func TestFloor(t *testing.T) {
	var Cases = []struct {
		number []string
		expect int
	}{
		{number: []string{"(", "(", ")", ")"}, expect: 0},
		{number: []string{"(", ")", "(", ")"}, expect: 0},
		{number: []string{"(", "(", "("}, expect: 3},
		{number: []string{"(", ")", ")"}, expect: 2},
	}
	for _, tc := range Cases {
		got := Floor(tc.number)
		if got!=tc.expect{
			t.Errorf("got %v\n expected %v",got,tc.expect)
		}
	}
}
