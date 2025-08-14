package iteration

import "testing"
import "reflect"

func TestSum(t *testing.T){
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("Want %d got %d ", want, got)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("Want %d got %d ", want, got)
		}
	})
}

func TestSumAll(t *testing.T){
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	got := SumAll([]int{1,2}, []int{1,2,3})
	want := []int{3, 6}

	checkSums(t, got ,want)
}