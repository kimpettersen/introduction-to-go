package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cats := []string{"Muscat", "Hopper", "Garfield", "Nop", "Mittens", "Stallman", "Snowball", "Queen Elizabeth", "Linus"}
	// range return two values when iterating a slice index and value
	for _, catName := range cats {
		meow(catName)
	}
}

func meow(catName string) {
	random := rand.Intn(3)
	time.Sleep(time.Duration(random) * time.Second)

	fmt.Printf("%s: Meow!!\n", catName)
}
