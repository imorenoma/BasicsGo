package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	num1 := os.Args[1]
	oper := os.Args[2]
	num2 := os.Args[3]

	number1, e := strconv.Atoi(num1)
	number2, e := strconv.Atoi(num2)

	if e == nil {
		//fmt.Println("Ans:", number1+number2)
		switch oper {
		case "+":
			fmt.Println(number1 + number2)
		case "-":
			fmt.Println(number1 - number2)
		case "multiplica":
			fmt.Println(number1 * number2)
		case "/":
			if number2 == 0 {
				fmt.Println("You try to divided by zero, try again")
			} else {
				fmt.Println(number1 / number2)
			}
		default:
			fmt.Println("operation not allowed, or not implemented")
		}
	}

}
