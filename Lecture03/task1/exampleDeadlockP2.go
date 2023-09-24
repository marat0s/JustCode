package main

func main() {
	// здесь созданы 2 канала, но ни один из них не может быть закрыт,
	// потому что они используются в 2 разных goroutine-ах
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() { // goroutine будет ждать пока в ch1 появится значение и отправит в ch2 значение 1
		<-ch1
		ch2 <- 1
	}()

	go func() { // goroutine будет ждать пока в ch2 появится значение и отправит в ch1 значение 1
		<-ch2
		ch1 <- 1
	}()

	select {}
}
