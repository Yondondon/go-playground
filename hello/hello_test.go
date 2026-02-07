package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("returns greeting with given name", func(t *testing.T) {
		got := Hello("Itadori", "english")
		want := "Hello, Itadori!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("returns default greeting if name isn't provided", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("returns greeting in italian", func(t *testing.T) {
		got := Hello("Itadori", "italian")
		want := "Ciao, Itadori!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("returns greeting in english if language isn't supported", func(t *testing.T) {
		got := Hello("Itadori", "aaaa")
		want := "Hello, Itadori!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
