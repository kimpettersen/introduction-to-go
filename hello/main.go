package main

import "fmt"

const cat = "cat"

func main() {

	// Declare a variable
	var name string

	// Assign a value to it
	name = "Muscat"

	// Infer the type of the variable with :=
	age := 3
	species := cat

	// %s means string
	// %d means integer
	// %v means default type
	fmt.Printf("Hello, my name is %s and I'm a %d year old %v\n", name, age, species)
}
