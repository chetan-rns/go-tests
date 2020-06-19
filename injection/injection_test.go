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
	buffer := bytes.Buffer{}
	Greet(&buffer, "test")

	got := buffer.String()
	want := "Hello test!"

	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}
