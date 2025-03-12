package main

import "fmt"

func main() {

	fmt.Println("1. Create a program that displays the message", "Hello world", "on the screen.\n")

	fmt.Println("Hello World\n")

	fmt.Println("2. Create a program that requests a number and then displays the message", "The number entered was [number].\n")

	var i int
	fmt.Scan(&i)
	fmt.Println("The number entered was", i, "\n")

	fmt.Println("3. Create a program that requests two numbers and prints the sum.\n")

	var a3, b3 float32

	fmt.Scan(&a3)
	fmt.Scan(&b3)

	fmt.Println("The sum of", a3, "+", b3, "is", a3+b3, "\n")

	fmt.Println("4. Create a program that requests the 4 numbers and displays the average.\n")

	var a4, b4, c4, d4 float32

	fmt.Scan(&a4)
	fmt.Scan(&b4)
	fmt.Scan(&c4)
	fmt.Scan(&d4)

	fmt.Println("The average of", a4, b4, c4, d4, "is", (a4+b4+c4+d4)/4, "\n")

	fmt.Println("5. Create a program that converts meters to centimeters.\n")

	var meters float32

	fmt.Scan(&meters)

	fmt.Println(meters, "m corresponds to", meters*100, "centimeters.\n")

}
