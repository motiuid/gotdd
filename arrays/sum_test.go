package arrays

import "testing"

import "reflect"

func TestSum(t *testing.T)  {
	t.Run("Collection of 5 numbers", func (t *testing.T)  {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T)  {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTail(t *testing.T)  {

	checkSum := func (t *testing.T, got, want []int)  {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}	
	}

	t.Run("Sum all the tails for an array except head or first element", 
	func (t *testing.T)  {
		got := SumAllTail([]int{1, 2}, []int{3, 9})
		want := []int{2, 9}
		checkSum(t, got, want)
	})
	t.Run("Safety sum empty slices", func (t *testing.T)  {
		got := SumAllTail([]int{}, []int{3, 9})
		want := []int{0, 9}
		checkSum(t, got, want)
	})
}