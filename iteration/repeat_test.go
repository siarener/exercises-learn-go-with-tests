package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("repeat character 'a' 0 times", func(t *testing.T) {
		got := Repeat("a", 0)
		want := ""

		if got != want {
			t.Errorf("expected %q but got %q", want, got)
		}
	})
	t.Run("repeat character 'a' 6 times", func(t *testing.T) {
		got := Repeat("a", 6)
		want := "aaaaaa"

		if got != want {
			t.Errorf("expected %q but got %q", want, got)
		}
	})
	t.Run("repeat character 'b' 4 times", func(t *testing.T) {
		got := Repeat("b", 4)
		want := "bbbb"

		if got != want {
			t.Errorf("expected %q but got %q", want, got)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeatedCharacters := Repeat("a", 5)
	fmt.Println(repeatedCharacters)
	// Output: aaaaa
}
