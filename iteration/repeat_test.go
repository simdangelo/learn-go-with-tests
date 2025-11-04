package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repetition by 3", func(t *testing.T){
		got := Repeat("a", 3)
		want := "aaa"
		if got != want {
			t.Errorf("expected %q got %q", want, got)
		}
	})
	t.Run("repetition by 10", func(t *testing.T) {
		got := Repeat("a", 10)
		want := "aaaaaaaaaa"
		if got != want {
			t.Errorf("expected %q got %q", want, got)
		}
	})
}

// go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 10)
	}
}