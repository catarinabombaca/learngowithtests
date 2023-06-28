package main

import "fmt"

const (
	french  = "French"
	spanish = "Spanish"

	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
)

func main() {
	fmt.Println(Hello("world", ""))
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}

	return
}
