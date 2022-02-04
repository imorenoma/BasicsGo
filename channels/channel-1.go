package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	/*---------
	Ejemplo1:
	-----------
	ch := make(chan int)
	wg.Add(2)

	//reciving go routine

	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	//send  go routine

	go func() {
		ch <- 42
		wg.Done()
	}()

	wg.Wait()*/

	/*-------------
	Ejemplo 2:
	--------------
	ch := make(chan int)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			ch <- 42
			wg.Done()
		}()
	}
	wg.Wait()*/

	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()
}
