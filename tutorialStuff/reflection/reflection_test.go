package reflection

import "testing"

func TestWa(t *testing.T) {
	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	Walk(x, func(input string) {
		got = append(got, input)
	})

	if got[0] != expected {
		t.Errorf("Wrong number of function calls, got %s want %s", got[0], expected)
	}
}
