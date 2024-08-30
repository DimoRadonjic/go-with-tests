package main

import "fmt"

const englishHelloPrefix = "Hello, "

const exlimationPoint = " !"

func HelloWorld(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name + exlimationPoint
}

func main() {
	fmt.Println(HelloWorld("Chris"))
}