package hello

import "testing"

func TestHello(t *testing.T) {

	t.Run("say hello to people", func(t *testing.T) {
		got := hello("Foo", "English")
		want := "Hello, Foo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is passed", func(t *testing.T) {
		got := hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("add prefix according to language", func(t *testing.T) {
		got := hello("Foo", "Kannada")
		want := "Namaskara, Foo"

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
