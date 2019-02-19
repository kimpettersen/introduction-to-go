package main

import "fmt"

func main() {
	for i := 0; i < 2; i++ {
		meow()
	}
}

func meow() {
	fmt.Println("Meow!!")
}
