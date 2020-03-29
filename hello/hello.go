package hello

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	kannadaHelloPrefix = "Namaskara, "
	spanishHelloPrefix = "Hola, "
)

func hello(name, language string) string {
	if name == "" {
		return "Hello, World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case "Kannada":
		prefix = kannadaHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(hello("World", "English"))
}