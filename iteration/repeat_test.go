package iteration

import "testing"
import "fmt"

func TestRepeating(t *testing.T){
	got := Repeating("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("No!")
	}
}

func BenchmarkRepeating(b *testing.B) {
	// Reset timer to exclude setup time from benchmark
	b.ResetTimer()

	// Loop b.N times for benchmarking
	for i := 0; i < b.N; i++ {
		Repeating("a", 12)
	}
}

func ExampleRepeat(){
	out := Repeating("a", 5)
	fmt.Println(out);
	// Output : "aaaaa"
}

func TestContainsA(t *testing.T){
	got := GotAnA("no way")
	want := true

	if got != want {
		t.Errorf("No!")
	}
}