package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 1 // это приведет к deadlock-у потому что нет другой goroutine чтобы считать из канала
	fmt.Println(<-ch)
}
