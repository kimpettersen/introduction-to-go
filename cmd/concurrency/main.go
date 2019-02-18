package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go meow(i, &wg)
	}
	wg.Wait()
}

func meow(catNum int, wg *sync.WaitGroup) {
	fmt.Printf("Cat #%d says: Meow!!\n", catNum)
	wg.Done()
}
