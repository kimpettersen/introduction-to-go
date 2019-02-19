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
	Animal
	Catnip bool
}

func (c Cat) LovesCatnip() bool {
	return c.Catnip
}

func main() {
	c := new(Cat)
	c.Name = "Muscat"
	c.Age = 3
	c.Catnip = true
	c.Says = "Meowwwww!"
	c.Speak()

	if c.LovesCatnip() {
		fmt.Printf("Party time for %s\n", c.Name)
	}
}
