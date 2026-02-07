package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeats character 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		assertCorrectMessage(t, got, want)
	})

	t.Run("returns empty string when count is 0", func(t *testing.T) {
		got := Repeat("a", 0)
		want := ""

		assertCorrectMessage(t, got, want)
	})

	t.Run("return empty string when count is less than 0", func(t *testing.T) {
		got := Repeat("a", -1)
		want := ""

		assertCorrectMessage(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
