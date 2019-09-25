package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrectSequence := func(t *testing.T, expected, repeated string) {
		t.Helper()
		if expected != repeated {
			t.Errorf("expected %q repeated %q", expected, repeated)
		}
	}

	t.Run("Repeats five times when specified", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectSequence(t, expected, repeated)
	})

	t.Run("Can pass in a count to repeat", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"
		assertCorrectSequence(t, expected, repeated)
	})
}
func ExampleRepeat() {
	repeated := Repeat("z", 3)
	fmt.Println(repeated)
	// Output: zzz
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
