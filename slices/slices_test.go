package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("array of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given, %v\n", got, want, numbers)
		}
	})
}

func Sum(numbers []int) int {
	sum := 0
	for _, a := range numbers {
		sum += a
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	sum := []int{}
	for _, slice := range slices {
		sum = append(sum, Sum(slice))
	}
	return sum
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	want := []int{3, 7}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTail(t *testing.T) {

	checkSums := func(t *testing.T, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum of slices", func(t *testing.T) {
		got := SumAllTail([]int{1, 2}, []int{3, 4, 5})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("sum of empty slices", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func SumAllTail(slices ...[]int) []int {
	sum := []int{}
	for _, slice := range slices {
		if len(slice) == 0 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, Sum(slice[1:]))
		}
	}
	return sum
}
