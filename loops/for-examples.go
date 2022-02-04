package main

import (
	"fmt"
)

func main() {
	/*------------------
		Example 1: Basico
	--------------------*/
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("------------------")
	/*------------------
		Example 2:
	--------------------*/

	for j, k := 0, 0; j < 5; j, k = j+1, k+1 {
		fmt.Println(j, k)
	}
	fmt.Println("------------------")
	/*------------------
		Example 3:
	--------------------*/

	for l := 0; l < 5; l++ {
		fmt.Println(l)
		if l%2 == 0 {
			l /= 2
		} else {
			l = 2*l + 1
		}
	}
	fmt.Println("------------------")
	/*------------------
		Example 4: otra forma de for
	--------------------*/
	m := 0

	for ; m < 5; m++ {
		fmt.Println(m)
	}
	fmt.Println(m)

	fmt.Println("------------------")
	/*------------------
		Example 4:
	--------------------*/

	n := 0

	for n < 5 {
		fmt.Println(n)
		n++
	}

	fmt.Println(n)

	fmt.Println("------------------")
	/*------------------
		Example 5:
	--------------------*/

	p := 0

	for {
		fmt.Println(p)
		p++
		if p == 5 {
			break
		}

	}

	fmt.Println("------------------")
	/*------------------
		Example 5:
	--------------------*/

	for q := 0; q < 10; q++ {
		if q%2 == 0 {
			continue
		}
		fmt.Println(q)
	}
	fmt.Println("------------------")
	/*------------------
		Example 6: For anidado
	--------------------*/
loop:
	for r := 1; r <= 3; r++ {
		for s := 1; s <= 3; s++ {
			fmt.Println(r * s)
			if r*s >= 3 {
				break loop
			}
		}
	}
	fmt.Println("------------------")
	/*------------------
		Example 7: Recorriendo array k,v
	--------------------*/
	t := []int{1, 2, 3}

	for key, value := range t {
		fmt.Println(key, value)
	}

	fmt.Println("------------------")
	/*------------------
		Example 8: Recorriendo un Map
	--------------------*/

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	for key, value := range statePopulations {
		fmt.Println(key, " ", value)
	}
	fmt.Println("\n")
	//Si queremos por ejemplo solo las claves

	for _, value := range statePopulations {
		fmt.Println(value)
	}
	fmt.Println("------------------")
	/*------------------
		Example 9:
	--------------------*/
	u := "hello Go!"

	for key, value := range u {
		fmt.Println(key, value)
	}

}
