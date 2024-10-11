package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("f", 3)
	fmt.Println(repeated)
	// Output: fff
}

func TestRepeat(t *testing.T) {
	t.Run("repeat 5", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertExpected(t, repeated, expected)
	})
	t.Run("repeat 1", func(t *testing.T) {
		repeated := Repeat("a", 1)
		expected := "a"
		assertExpected(t, repeated, expected)
	})
}

func assertExpected(t testing.TB, repeated string, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q, got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
