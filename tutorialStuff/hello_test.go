package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Jan", "")
		want := "Hello, Jan"

		assertCorrectMessage(t, got, want);
	})

	t.Run("Saying hello to people in english", func(t *testing.T) {
		got := Hello("Jan", "English")
		want := "Hello, Jan"

		assertCorrectMessage(t, got, want);
	})

	t.Run("Saying hello to people in spanish", func(t *testing.T) {
		got := Hello("Jan", "Spanish")
		want := "Hola, Jan"

		assertCorrectMessage(t, got, want);
	})

	t.Run("Saying hello to people in french", func(t *testing.T) {
		got := Hello("Jan", "French")
		want := "Baguette, Jan"

		assertCorrectMessage(t, got, want);
	})

	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want);
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}