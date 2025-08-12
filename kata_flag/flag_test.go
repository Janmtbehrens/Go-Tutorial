package kataflag

import (
	"reflect"
	"testing"
)

func TestFlag(t *testing.T) {
	allFlags := "-d -args what,yes -timestamp 100 -p ./flag_test.go -v"

	t.Run("Test Debug Flag", func(t *testing.T) {
		want := true
		got := FindFlag(GetSchema(), "d").Evaluate(allFlags).Bool

		assertCorrectFlag(t, got, want)
	})
	t.Run("Test Path", func(t *testing.T) {
		want := "./flag_test.go"
		got := FindFlag(GetSchema(), "p").Evaluate(allFlags).String

		if want != got {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("Test Args Inputs", func(t *testing.T) {
		want := []string{"what", "yes"}
		got := FindFlag(GetSchema(), "args").Evaluate(allFlags).Array

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("Test Verbose Inputs", func(t *testing.T) {
		want := true
		got := FindFlag(GetSchema(), "v").Evaluate(allFlags).Bool

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %t, want %t", got, want)
		}
	})
	t.Run("Test Timestamp Inputs", func(t *testing.T) {
		want := 100

		got := FindFlag(GetSchema(), "timestamp").Evaluate(allFlags).Integer

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
	t.Run("Test Negative Timestamp Inputs", func(t *testing.T) {
		want := -100

		got := FindFlag(GetSchema(), "timestamp").Evaluate("-d -timestamp -100 -v").Integer

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
	noFlags := ""
	t.Run("Test No Debug Flag", func(t *testing.T) {
		want := false
		got := FindFlag(GetSchema(), "d").Evaluate(noFlags).Bool

		assertCorrectFlag(t, got, want)
	})
	t.Run("Test Default Path", func(t *testing.T) {
		want := "."
		got := FindFlag(GetSchema(), "p").Evaluate(noFlags).String

		if want != got {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("Test No Args Inputs", func(t *testing.T) {
		var want []string
		got := FindFlag(GetSchema(), "args").Evaluate(noFlags).Array

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("Test No Verbose Inputs", func(t *testing.T) {
		want := false

		got := FindFlag(GetSchema(), "v").Evaluate(noFlags).Bool

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %t, want %t", got, want)
		}
	})
	t.Run("Test No Timestamp Inputs", func(t *testing.T) {
		want := 1

		got := FindFlag(GetSchema(), "timestamp").Evaluate(noFlags).Integer

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
}

func assertCorrectFlag(t *testing.T, got, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("Got %t, want %t", got, want)
	}
}
