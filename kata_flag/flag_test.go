package kataflag

import (
	"reflect"
	"testing"
)

func TestFlag(t *testing.T) {
	allFlags := "-d -args what,yes -p ./flag_test.go -v"

	t.Run("Test Debug Flag", func(t *testing.T) {
		want := true
		got := InDebugMode(GetSchema(), allFlags)

		assertCorrectFlag(t, got, want)
	})
	t.Run("Test Path", func(t *testing.T) {
		want := "./flag_test.go"
		got := GetPath(GetSchema(), allFlags)

		if want != got {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("Test Args Inputs", func(t *testing.T) {
		want := []string{"what", "yes"}
		got := GetArgs(GetSchema(), allFlags)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("Test Verbose Inputs", func(t *testing.T) {
		want := true
		got := IsVerbose(GetSchema(), allFlags)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %t, want %t", got, want)
		}
	})
	noFlags := ""
	t.Run("Test No Debug Flag", func(t *testing.T) {
		want := false
		got := InDebugMode(GetSchema(), noFlags)

		assertCorrectFlag(t, got, want)
	})
	t.Run("Test No Path", func(t *testing.T) {
		want := ""
		got := GetPath(GetSchema(), noFlags)

		if want != got {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("Test No Args Inputs", func(t *testing.T) {
		want := []string{""}
		got := GetArgs(GetSchema(), noFlags)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("Test No Verbose Inputs", func(t *testing.T) {
		want := false

		got := IsVerbose(GetSchema(), noFlags)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %t, want %t", got, want)
		}
	})
}

func assertCorrectFlag(t *testing.T, got, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("Got %t, want %t", got, want)
	}
}
