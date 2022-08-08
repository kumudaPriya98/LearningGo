package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("Test the Sum for Array of size 5", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum5(numbers)
		want := 15

		if got != want {
			t.Errorf("got '%d' want '%d', Nums : %v", got, want, numbers)
		}
	})

	t.Run("Test the Sum for Slices", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

		got := Sum(numbers)
		want := 36

		if got != want {
			t.Errorf("got '%d' want '%d', Nums : %v", got, want, numbers)
		}
	})

}
