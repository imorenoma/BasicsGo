package main

import "fmt"

func main() {
	var (
		name   string  = "Jose"
		edad   int     = 11
		altura float32 = 1.20
	)

	var (
		alimento    string = "banana"
		lugar       string = "canarias"
		trackNumber int    = 123456789
	)

	fmt.Printf("%s, %v, %.2f\n", name, edad, altura)
	fmt.Printf("%s, %s, %v", alimento, lugar, trackNumber)
}
