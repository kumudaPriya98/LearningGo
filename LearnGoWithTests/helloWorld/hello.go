package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const helloPrefix = "Hello "
const spanishPrefix = "Hola "
const frenchPrefix = "Bonjour "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := helloPrefix
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Kumuda", ""))
}
