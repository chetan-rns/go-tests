package injection

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello %s!", name)
}

func TestGreet(t *testing.T) {
	buffer := &bytes.Buffer{}
	Greet(buffer, "Foo")

	got := buffer.String()
	want := "Hello Foo!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
