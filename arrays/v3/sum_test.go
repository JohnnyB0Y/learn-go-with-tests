package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("collections of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3}
		slice := numbers[0:]

		got := Sum(numbers)
		got2 := SumGo(slice)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

		if got2 != want {
			t.Errorf("got2 %d want %d given, %v", got2, want, numbers)
		}
	})

}
