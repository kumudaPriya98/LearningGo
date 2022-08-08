package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	checkEqualSum := func(t testing.TB, got, want int) {
		t.Helper()

		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	}

	t.Run("Test the Sum for Array of size 5", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum5(numbers)
		want := 15
		checkEqualSum(t, got, want)
	})

	t.Run("Test the Sum for Slices", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

		got := Sum(numbers)
		want := 36
		checkEqualSum(t, got, want)
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3, 4}, []int{6, 7})
	want := []int{10, 13}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v are not in equal size", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2, 3, 4}, []int{6, 7}, []int{9}, []int{})
	want := []int{9, 7, 0, 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v are not same", got, want)
	}
}