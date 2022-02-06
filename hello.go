package main

import (
	"fmt"
)

// Hello hello信息
func Hello(name string) string {
	return "Hello," + name + "!"
}

func main() {
	fmt.Printf(Hello("world"))
}
