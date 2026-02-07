package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum([]int{1, 2, 3})
	want := 6

	if got != want {
		t.Errorf("got '%d', want '%d'", got, want)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{2, 2})
	want := []int{6, 4}

	if !slices.Equal(got, want) {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}

func BenchmarkSumAll(b *testing.B) {
	for b.Loop() {
		SumAll([]int{1, 2, 3}, []int{2, 2}, []int{1, 4}, []int{3, 2})
	}
}
