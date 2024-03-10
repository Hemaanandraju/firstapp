package main

import (
	"fmt"
)

func main() {
	egVAR()
	shadowing()
	firstapp()
	secondapp()
	booltest()
	stringtest()

}

func egVAR() {
	var (
		A = 10.00
		B = 20
		C = float32(10.01)
	)
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

func booltest() {
	var n = 1 != 1
	fmt.Printf("%v, %T", n, n)
}

func stringtest() {
	var s string
	s = "this is string"
	fmt.Println("\n", s)
}
