package main

import "fmt"

const englishPrefix = "Hello, "

func main() {
	fmt.Println(Hello("world"))
}

func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return englishPrefix + name
}
