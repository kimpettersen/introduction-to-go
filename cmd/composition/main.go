package main

import "fmt"

type Animal struct {
	Name string
	Age  int
	Says string
}

func (a Animal) Speak() {
	if a.Says == "" {
		fmt.Println("Howdy!")
		return
	}
	fmt.Println(a.Says)
}

type Cat struct {
	// Go does not support inheritance
	// You can Compose a struct out of other structs. Here we say that a Cat is an Animal
	Animal

	CatnipLover bool
}

func (c Cat) LovesCatnip() bool {
	return c.CatnipLover
}

func main() {
	c := new(Cat)
	c.Name = "Muscat"
	c.Age = 3
	c.CatnipLover = true
	c.Says = "Meowwwww!"

	// Since Cat is an Animal, we can access the speak method
	c.Speak()

	// And, of course, the LovesCatnip method
	if c.LovesCatnip() {
		fmt.Printf("Party time for %s\n", c.Name)
	}
}
