package main

import "fmt"

// Hello returns a personalised greeting.
func Hello(name string) string {
	return "Hello, " + name
}

func HelloGo(name string) string {
	return "Hello Go, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
