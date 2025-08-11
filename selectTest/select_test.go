package selectTest

import "testing"

func TestRacer(t *testing.T) {
	slow := "http://www.facebook.com"
	fast := "http://www.quii.dev"

	want := fast
	got := Racer(slow, fast)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

