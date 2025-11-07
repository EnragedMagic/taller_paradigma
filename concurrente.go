package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)
// Esta es la estructura de la tarea 
type Tarea struct {
	id        int
	prioridad int
	duracion  time.Duration
}
// Esta ejecuta la tarea y reporta el resultado 
func ejecutarTarea(t Tarea, resultados chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(t.duracion)
	resultados <- fmt.Sprintf("Tarea %d con prioridad %d completada", t.id, t.prioridad)
}

func main() {
	rand.Seed(time.Now().UnixNano())
// Lista de tareas 
	tareas := []Tarea{
		{1, 1, time.Duration(rand.Intn(1000)) * time.Millisecond},
		{2, 3, time.Duration(rand.Intn(1000)) * time.Millisecond},
		{3, 2, time.Duration(rand.Intn(1000)) * time.Millisecond},
		{4, 1, time.Duration(rand.Intn(1000)) * time.Millisecond},
		{5, 2, time.Duration(rand.Intn(1000)) * time.Millisecond},
	}

	sort.Slice(tareas, func(i, j int) bool {
		return tareas[i].prioridad < tareas[j].prioridad
	})

	canalTareas := make(chan Tarea)
	resultados := make(chan string)
	var wg sync.WaitGroup

	go func() {
		for _, t := range tareas {
			canalTareas <- t
		}
		close(canalTareas)
	}()

	for i := 0; i < 3; i++ {
		go func() {
			for t := range canalTareas {
				wg.Add(1)
				go ejecutarTarea(t, resultados, &wg)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultados)
	}()

	for res := range resultados {
		fmt.Println(res)
	}

	fmt.Println("Todas las tareas han sido completadas.")
}
