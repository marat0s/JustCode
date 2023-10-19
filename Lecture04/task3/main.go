package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var rwMutex sync.RWMutex

func readData(id int) {
	rwMutex.RLock()
	fmt.Printf("Reader %d: Count = %d\n", id, count)
	time.Sleep(1 * time.Second) // симулируем чтение
	rwMutex.RUnlock()
}

func writeData(id int) {
	rwMutex.Lock()
	count++
	fmt.Printf("Writer %d updated count to %d\n", id, count)
	rwMutex.Unlock()
}

func main() {
	go readData(1)  // Reader 1
	go readData(2)  // Reader 2
	go writeData(1) // Writer 1
	go writeData(2) // Writer 2

	time.Sleep(5 * time.Second) // ждем когда все goroutine-ы завершатся
}
