package main

import (
	"fmt"
	"sync"
	"time"
)
// Funcion 1
func imprimirNumeros(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println("Numero:", i)
		time.Sleep(300 * time.Millisecond)
	}
}
// Funcion 2 
func imprimirLetras(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, letra := range []string{"A", "B", "C", "D", "E"} {
		fmt.Println("Letra:", letra)
		time.Sleep(400 * time.Millisecond)
	}
}
// Funcion 3
func mostrarEstado(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("Procesando...")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go imprimirNumeros(&wg)
	go imprimirLetras(&wg)
	go mostrarEstado(&wg)

	wg.Wait()
	fmt.Println("Todas las goroutines han finalizado.")
}
