package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type Mouse struct {
	Name  string
	Age   int
	Hates string
}

func (m Mouse) IsHating(name string) (bool, error) {
	if name == "" {
		return false, errors.New("Error: name cannot be empty")
	}
	return m.Hates != name, nil
}

func main() {
	m := Mouse{
		Name:  "Jerry",
		Age:   3,
		Hates: "Tom",
	}

	// Will this end in an error?
	_, err := m.IsHating("")
	if err != nil {
		fmt.Println(err)
	}

	// Hate gone strong
	hate, err := m.IsHating("Tom")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s hates Tom: %t\n", m.Name, hate)
}
