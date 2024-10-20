package main

import (
	"fmt"
)

const greetingPrefix = "Hello, "

func Hello(who string) string {
	if who == "" {
		who = "World"
	}
	return greetingPrefix + who
}

func main() {
	fmt.Println(Hello(""))
}
