package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func() {
		escrever("Olá mundo") //goroutine
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em GO")
		waitGroup.Done()
	}()

	go func() {
		escrever("goroutine 3") //goroutine
		waitGroup.Done()
	}()

	go func() {
		escrever("goroutine 4")
		waitGroup.Done()
	}()

	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
