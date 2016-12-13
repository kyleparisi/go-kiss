package main

import "fmt"

// reference: https://nathanleclaire.com/blog/2015/10/10/interfaces-and-composition-for-effective-unit-testing-in-golang/

func hello() string {
	return "Hello, Testing!"
}

func main() {
	fmt.Println(hello())
}
