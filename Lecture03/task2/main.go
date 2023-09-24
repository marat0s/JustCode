package main

import (
	"fmt"
)

func mergeChannels(channels ...chan int) chan int { // функция принимает неопределенное количество каналов
	merged := make(chan int)
	go func() {
		defer close(merged)           // вызываем в конце функции
		for _, ch := range channels { // проходимся по всем каналам
			for val := range ch { // читаем из канала пока он не закроется
				merged <- val
			}
		}
	}()
	return merged
}

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)

	ch1 <- 1
	ch1 <- 2
	close(ch1)

	ch2 <- 3
	ch2 <- 4
	close(ch2)

	merged := mergeChannels(ch1, ch2)
	for val := range merged {
		fmt.Println(val)
	}
}
