package main

import (
	"fmt"
)
// Esta funcion produce
func productora(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}
// Esta funcion consume
func consumidora(ch chan int, done chan bool) {
	for num := range ch {
		if num%2 == 0 {
			fmt.Println("Numero par:", num)
		}
	}
	done <- true
}
// Funcion main
func main() {
	ch := make(chan int)
	done := make(chan bool)

	go productora(ch)
	go consumidora(ch, done)

	<-done
	fmt.Println("Fin del procesamiento de canales.")
}
