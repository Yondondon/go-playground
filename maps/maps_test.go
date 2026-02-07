package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known key", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		if got == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new key", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is just a test"
		dictionary.Add(key, value)

		assertDefinition(t, dictionary, key, value)
	})

	t.Run("existing key", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, "new test")

		assertError(t, err, ErrKeyExists)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing key", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{key: value}
		newValue := "new definition"

		err := dictionary.Update(key, newValue)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, newValue)
	})

	t.Run("new key", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(key, value)

		assertError(t, err, ErrKeyNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing key", func(t *testing.T) {
		key := "test"
		value := "test value"
		dictionary := Dictionary{key: value}

		dictionary.Delete(key)
		_, err := dictionary.Search(key)

		assertError(t, err, ErrNotFound)
	})

	t.Run("non existing key", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Delete("nonExistingKey")
		assertError(t, err, ErrKeyNotFound)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, definition string) {
	t.Helper()

	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added key:", err)
	}
	assertStrings(t, got, definition)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
