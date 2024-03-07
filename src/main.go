package main

import (
	"fmt"
)

func main() {
	egVAR()
	shadowing()
	firstapp()
	secondapp()

}

func egVAR() {

	var A int
	A = 10.00
	B := 20
	var C = float32(10.01)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}

func shadowing() {
	// Outer variable
	x := 10
	fmt.Println("Outer x:", x) // Output: Outer x: 10

	// Inner scope with shadowing
	{
		// Inner variable shadowing the outer variable
		x := 20
		fmt.Println("Inner x:", x) // Output: Inner x: 20
	}

	// Accessing the outer variable again
	fmt.Println("Outer x after inner scope:", x) // Output: Outer x after inner scope: 10
}

func firstapp() {
	fmt.Println("My First App")
}
