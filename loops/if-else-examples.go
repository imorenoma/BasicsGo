package main

import (
	"fmt"
	"math"
)

func main() {

	/*------------------
		Example 1:
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

	fmt.Println("\n")

	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println("Example 1")
		fmt.Println(pop, "\n")
	}

	/*------------------
		Example 2:
	--------------------*/

	number := 50
	guess := 30

	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	}

	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("Example 2")
			fmt.Println("Too low", "\n")
		}

		if guess > number {
			fmt.Println("Example 2")
			fmt.Println("Too high", "\n")
		}

		if guess == number {
			fmt.Println("Example 2")
			fmt.Println("You got it", "\n")
		}

		fmt.Println(number <= guess, number >= guess, number != guess, "\n")
	}

	/*------------------
		Example 3:
	--------------------*/

	number2 := 50
	guess2 := 60

	if guess2 < 1 || guess2 > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	} else {
		if guess2 < number2 {
			fmt.Println("Example 3")
			fmt.Println("Too low", "\n")
		}

		if guess2 > number2 {
			fmt.Println("Example 3")
			fmt.Println("Too high", "\n")
		}

		if guess2 == number2 {
			fmt.Println("Example 3")
			fmt.Println("You got it", "\n")
		}

		fmt.Println(number2 <= guess2, number2 >= guess2, number2 != guess2, "\n")
	}

	/*------------------
		Example 4:
	--------------------*/
	number3 := 50
	guess3 := 101

	if guess3 < 1 || guess3 > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	} else if guess3 > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {
		if guess3 < number3 {
			fmt.Println("Too low")
		}
		if guess3 > number3 {
			fmt.Println("Too high")
		}
		if guess3 == number3 {
			fmt.Println("You gor it!")
		}

		fmt.Println(number3 <= guess3, number3 >= guess3, number3 != guess3)
	}

	/*------------------
		Example 5:
	--------------------*/

	myNum := 0.123

	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("This are different")
	}

}
