package main

import "fmt"

func main() {
	cats := []string{"Muscat", "Tom", "Mons", ""}
	for _, catName := range cats {
		meow(catName)
	}
}

func meow(catName string) {
	fmt.Printf("%s: Meow!!\n", catName)
}
