package main

import "fmt"

func main() {

	fmt.Println("1. Create a program that displays the message", "Hello world", "on the screen.\n")

	fmt.Println("Hello World\n")

	fmt.Println("2. Create a program that requests a number and then displays the message", "The number entered was [number].\n")

	var i int
	fmt.Scan(&i)
	fmt.Println("The number entered was", i)

	fmt.Println("3. Create a program that requests two numbers and prints the sum.")

	var a, b float32

	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println("The sum of", a, "+", b, "is", a+b)
}
