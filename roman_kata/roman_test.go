package romankata

import (
	"reflect"
	"testing"
)

func TestDigitToRoman(t *testing.T) {
	t.Run("Convert random numbers", func(t *testing.T) {
		arabicNumbers := []int{1, 12, 54, 127, 145}

		want := []Roman{"I", "XII", "LIV", "CXXVII", "CXLIV", "MCLIV"}
		got := make([]Roman, len(arabicNumbers))

		for _, num := range arabicNumbers {
			got = append(got, ToRoman(num))
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("Convert first 10 numbers", func(t *testing.T) {
		arabicNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		want := []Roman{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
		got := make([]Roman, len(arabicNumbers))

		for _, num := range arabicNumbers {
			got = append(got, ToRoman(num))
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
}
