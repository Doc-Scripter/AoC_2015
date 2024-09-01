package main

import (
	"testing"
)

func TestDeliver(t *testing.T) {
	Testcases := []struct {
		input string
		want  int
	}{
		{input: "^v", want: 3},
		{input: "^>v<", want: 3},
	}

	for _, test := range Testcases {
		got := Deliver(test.input)
		if got != test.want {
			t.Errorf("got %v\n expected %v\n", got, test.want)
		}

	}
}
