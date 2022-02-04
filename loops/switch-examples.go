package main

import "fmt"

func main() {

	/*------------------
		Example 1:
	--------------------*/

	switch 2 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	/*------------------
		Example 2:
	--------------------*/

	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	/*------------------
		Example 3:
	--------------------*/
	j := 10

	switch {
	case j <= 10:
		fmt.Println("less than or equal to ten")
		//fallthrough esta sentecia seria como pass de python
	case j <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	/*------------------
		Example 4:
	--------------------*/

	var k interface{} = [3]int{}

	switch k.(type) {
	case int:
		fmt.Println("k is an int")
	case float64:
		fmt.Println("k is an float64")
	case string:
		fmt.Println("k is string")
	case [3]int:
		fmt.Println("k is [3]int")
	default:
		fmt.Println("k is anoter type")
	}
}
