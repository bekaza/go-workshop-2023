package main

import (
	"fmt"
)

func Hello() string {
	return "Hello, world 2023"
}

func HelloWithMessage(msg string) string {
	if msg != "test" {
		return Hello()
	}
	return fmt.Sprintf("%s - 2023", msg)
}

func main() {
	fmt.Println(Hello())
	msg := "Golang Workshop"
	fmt.Println(HelloWithMessage(msg))
}
