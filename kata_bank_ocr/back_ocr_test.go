package katabankocr

import "testing"

func TestReader(t *testing.T) {
	t.Run("Test chain", func(t *testing.T) {
		want := "123456789"
		got, _ := Read("./bankNumbers.txt", 0)

		if want != got {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	/*
		t.Run("051", func(t *testing.T) {
			want := "0000000051"
			got, _ := Read("./bankNumbers.txt", 1)

			if want != got {
				t.Errorf("Got %s, want %s", got, want)
			}
		})
		t.Run("4 Chain", func(t *testing.T) {
			want := "444444444"
			got, _ := Read("./bankNumbers.txt", 2)

			if want != got {
				t.Errorf("Got %s, want %s", got, want)
			}
		})
		t.Run("Borken", func(t *testing.T) {
			want := "1234?678? ILL"
			got, _ := Read("./bankNumbers.txt", 3)

			if want != got {
				t.Errorf("Got %s, want %s", got, want)
			}
		})
		t.Run("Ambiguous", func(t *testing.T) {
			want := []string{"666666666", "666566666", "686666666"}
			got, _ := Read("./bankNumbers.txt", 4)

			if want != got {
				t.Errorf("Got %s, want %s", got, want)
			}
		})
		t.Run("Too many", func(t *testing.T) {
			want := "1234567899"
			got, err := Read("./bankNumbers.txt", 5)

			if err == nil {
				t.Errorf("Wanted error but got none")
			}
			if want != got {
				t.Errorf("Got %s, want %s", got, want)
			}
		})
		t.Run("Too few", func(t *testing.T) {
			want := "12345678"
			got, err := Read("./bankNumbers.txt", 6)

			if err == nil {
				t.Errorf("Wanted error but got none")
			}
			if want != got {
				t.Errorf("Got %s, want %s", got, want)
			}
		})*/
}
