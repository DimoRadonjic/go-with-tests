package main

import "fmt"

// const englishHelloPrefix = "Hello, "
// const spanishHelloPrefix = "Hola, "
// const frenchHelloPrefix = "Bonjour, "
// const spanish = "Spanish"
// const french = "French"
// const exlimationPoint = " !"

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	exlimationPoint = " !"
)

func greetingPrefix(lang string) (prefix string)  {
	switch lang {
		case spanish:
			prefix = spanishHelloPrefix 
		case french : 
			prefix = frenchHelloPrefix 
		default : 
			prefix = englishHelloPrefix 
	}
	return
}

func HelloWorld(name, lang string) string {
	if name == "" {
		name = "World"
	}

	// if lang == spanish {
	// 	return spanishHelloPrefix + name + exlimationPoint
	// }
	// if lang == french{
	// 	return frenchHelloPrefix + name + exlimationPoint
	// }
	// return englishHelloPrefix + name + exlimationPoint
	return greetingPrefix(lang) + name + exlimationPoint
}

func main() {
	fmt.Println(HelloWorld("Napoleon", "German"))
}