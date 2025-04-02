package helloworld

import "fmt"

type Language int

const (
	English Language = iota
	Spanish
	French
)

const (
	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Hola"
	frenchHelloPrefix  = "Bonjour"
)

// Print hello to user depending on the language
func Hello(name string, lang Language) string {
	if name == "" {
		name = "world"
	}

	return fmt.Sprintf("%s, %s", greetingPrefix(lang), name)
}

func greetingPrefix(lang Language) (prefix string) {
	switch lang {
	case English:
		prefix = englishHelloPrefix
	case Spanish:
		prefix = spanishHelloPrefix
	case French:
		prefix = frenchHelloPrefix
	}

	return
}
