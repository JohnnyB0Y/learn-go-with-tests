package main

import "fmt"

const englishHelloPrefix = "Hello, "
const chineseHelloPrefix = "泥猴，"

// Hello returns a personalised greeting.
func Hello(name string) string {
	return englishHelloPrefix + name
}

func HelloGo(name string) string {
	return chineseHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
