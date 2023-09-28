package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	fmt.Println(counter) // так как counter не является потокобезопасным, то мы можем получить любое значение от 0 до 1000
}
