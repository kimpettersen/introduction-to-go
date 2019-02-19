package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	cats := []string{"Muscat", "Tom", "Mons", ""}
	for _, catName := range cats {
		wg.Add(1)
		go meow(catName, &wg)
	}
	wg.Wait()
}

func meow(catName string, wg *sync.WaitGroup) {
	fmt.Printf("%s: Meow!!\n", catName)
	wg.Done()
}
