package main

import "fmt"

// Hello returns a greeting.
func Hello() string {
	return "Hello, world"
}

func HelloGo() string {
	return "Hello Go!"
}

func main() {
	fmt.Println(Hello())
}
