package integer

import "testing"
import "fmt"

func TestAdd(t *testing.T) {
	got := Add(2,1)
	want := 3

	assertCorrectMessageInt(t, got, want)
}

func assertCorrectMessageInt(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}