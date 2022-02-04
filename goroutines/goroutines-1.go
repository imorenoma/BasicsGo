package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello() //go es una abstraccion de hilo
	time.Sleep(100 * time.Microsecond)

}

func sayHello() {
	fmt.Println("Hello")
}
