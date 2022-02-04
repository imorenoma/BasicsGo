package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{} //mutex no es mas que un candado

func main() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}

	wg.Wait()
}

func sayHello() {

	fmt.Printf("Hello # %v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

/*
func main() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}

	wg.Wait()

}

func sayHello() {
	m.RLock() //tipo read lock
	fmt.Printf("Hello # %v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	m.Lock() //tipo write lock
	counter++
	m.Unlock()
	wg.Done()
}
*/
