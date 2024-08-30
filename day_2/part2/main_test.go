package main

import "testing"

func TestTotalArea(t *testing.T) {
	test := []struct {
		input  []int
		expect int
	}{
		{input:[]int{2,3,4},expect: 58},
	}
	for _,tc:=range test{
		got:=TotalArea(tc.input)
		if got!=tc.expect{
			t.Errorf("got :%v \n expect : %v\n",got,tc.expect)
		}
	}
}
