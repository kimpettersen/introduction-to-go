package main

import "fmt"

type Cat struct {
	Name   string
	Age    int
	Catnip bool
}

func (c Cat) LovesCatnip() bool {
	return c.Catnip
}

func main() {
	c := Cat{
		Name:   "Muscat",
		Age:    3,
		Catnip: true,
	}

	if c.LovesCatnip() {
		fmt.Printf("Party time for %s\n", c.Name)
	}
}
