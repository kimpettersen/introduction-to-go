package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	cats := []string{"Muscat", "Hopper", "Garfield", "Nop", "Mittens", "Stallman", "Snowball", "Queen Elizabeth", "Linus"}
	for _, catName := range cats {
		wg.Add(1)
		go meow(catName, &wg)
	}
	wg.Wait()
}

func meow(catName string, wg *sync.WaitGroup) {
	// Sleep for random seconds
	random := rand.Intn(3)
	time.Sleep(time.Duration(random) * time.Second)

	fmt.Printf("%s: Meow!!\n", catName)
	wg.Done()
}
