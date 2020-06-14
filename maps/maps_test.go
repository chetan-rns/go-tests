package maps

import (
	"errors"
	"testing"
)

var (
	errKeyNotFound      = errors.New("key not found")
	errKeyAlreadyExists = errors.New("key already exists")
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", errKeyNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, ok := d[key]
	if ok {
		return errKeyAlreadyExists
	}
	d[key] = value
	return nil
}

func TestSearch(t *testing.T) {
	t.Run("key found in dictionary", func(t *testing.T) {
		dict := Dictionary{"test": "this is testing"}

		want := "this is testing"
		got, err := dict.Search("test")
		assertError(t, err, nil)
		assertString(t, got, want)
	})
	t.Run("key not found", func(t *testing.T) {
		dict := Dictionary{"test": "this is testing"}

		want := errKeyNotFound
		_, err := dict.Search("fake")

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add a new key:value", func(t *testing.T) {
		dict := Dictionary{"2": "4"}
		err := dict.Add("3", "9")
		assertNoError(t, err)

		want := "9"
		got, err := dict.Search("3")

		assertNoError(t, err)
		assertString(t, got, want)
	})
	t.Run("key already exists", func(t *testing.T) {
		dict := Dictionary{"key": "value"}

		want := errKeyAlreadyExists
		err := dict.Add("key", "value")

		assertError(t, err, want)
	})
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if !errors.Is(want, got) {
		t.Fatalf("got %v want %v", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
