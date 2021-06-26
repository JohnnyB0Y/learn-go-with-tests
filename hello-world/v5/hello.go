package main

import "fmt"

const englishHelloPrefix = "Hello, "
const chineseHelloPrefix = "泥猴，"

// Hello returns a personalised greeting, defaulting to Hello, world if an empty name is passed.
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func HelloGo(name string) string {
	if name == "" {
		name = "World"
	}
	return chineseHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
