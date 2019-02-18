package main

import "fmt"

const cat = "cat"

func main() {
	var name string
	name = "Muscat"
	age := 3
	species := cat

	fmt.Printf("Hello, my name is %s and I'm a %d year old %v\n", name, age, species)
}
