package main

import (
	"fmt"
	"time"
)

func main() {
	//playing with channels
	ch := make(chan int)

	go send(ch)
	receive(ch)
}
func receive(ch chan int) {
	for i := 0; i < 10; i++ {

		fmt.Println(<-ch)

	}

}
func send(ch chan int) {
	for i := 0; i < 10; i++ {

		ch <- i
		time.Sleep(100 * time.Millisecond)

	}

}
