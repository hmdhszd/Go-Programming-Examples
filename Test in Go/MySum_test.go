package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{1, 2, 3, 4, 5, 5, 22}, 42},
		test{[]int{1, 1}, 2},
		test{[]int{-1, 0, 1}, 0},
		test{[]int{3, 4}, 7},
		test{[]int{333, 333}, 666},
	}

	for _, v := range tests {

		output := MySum(v.data...)
		if output != v.answer {
			t.Error("Error, Expected", v.answer, "but got", output)
		}

	}

}
