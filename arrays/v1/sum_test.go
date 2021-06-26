package main

import "testing"

func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSumGo(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := SumGo(numbers)
	want := 15

	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

var numbers = [5]int{1, 2, 3, 4, 5}

func BenchmarkSum(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Sum(numbers)
	}
}

func BenchmarkSumGo(t *testing.B) {
	for i := 0; i < t.N; i++ {
		SumGo(numbers)
	}
}
