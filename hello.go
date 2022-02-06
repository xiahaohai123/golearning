package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const englishHelloSuffix = "!"

// Hello hello信息
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name + englishHelloSuffix
}

func main() {
	fmt.Printf(Hello("world"))
}
