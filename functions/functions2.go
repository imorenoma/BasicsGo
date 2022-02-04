package main

import "fmt"

func main() {
	/*
		sayMessage("Hello Go!")

		/*------------------------------
			Example 1: message function
		------------------------------

		for i := 0; i < 5; i++ {
			message("Hello world !", i)
		}
		/*------------------------------
			Example 2: same type param.
		------------------------------

		sayGreating("Hello", "Susan")
		/*--------------------------------------
			Example 3: sustitution with pointers
		---------------------------------------
		greating := "Hello"
		name := "Stacy"
		newGreating(&greating, &name)
		fmt.Println(name)
		/*--------------------------------------
			Example 4: variadic parameters
		---------------------------------------

		sum(1, 2, 3, 4, 5)

		/*--------------------------------------
			Example 5:funtion with return
		---------------------------------------

		s := secondSum(1, 2, 3, 4, 5)
		fmt.Println("the sum is:", s)

		/*----------------------------------------
			Example 6: funtion with multiple return
		-----------------------------------------

		d, err := divided(5.0, 0.0)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(d)

		/*----------------------------------------
			Example 7: Anonimous Function
		-----------------------------------------

		func() {
			fmt.Println("Hi Go!!!")
		}()

		for j := 0; j < 5; j++ {
			func(i int) {
				fmt.Println(j)
			}(j)
		}

		/*----------------------------------------
			Example 8: Method
		-----------------------------------------*/

	g := greeter{
		saludo: "Hello",
		nombre: "Ayaka",
	}
	g.greet()
	fmt.Println("the new name is", g.nombre)

}

/*
func sayMessage(msg string) {
	fmt.Println(msg)
}

func message(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("the value of index is", idx)
}

//Just in case we have same type of parameters Go
//Use sintactic sugar and you can write a functions as
//bellow

func sayGreating(greeting, name string) {
	fmt.Println(greeting, name)
}

func newGreating(greating, name *string) {
	fmt.Println(*greating, *name)
	*name = "Ted"
	fmt.Println(*name)
}

func sum(values ...int) {
	fmt.Println(values)
	result := 0

	for _, v := range values {
		result += v
	}
	fmt.Println("the sum is", result)
}

func secondSum(values ...int) int {
	fmt.Println(values)
	result := 0

	for _, v := range values {
		result += v
	}
	return result
}

func divided(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot diveded by 0")
	}
	return a / b, nil

}*/

type greeter struct {
	saludo string
	nombre string
}

func (g greeter) greet() {
	fmt.Println(g.saludo, g.nombre)
	g.nombre = ""
}
