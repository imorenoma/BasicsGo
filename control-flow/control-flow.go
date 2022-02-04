package main

import "fmt"

func main() {

	/*------------------
		Example 1: Basico defer
	--------------------*/        //  / \
	fmt.Println("init")           // / | \
	defer fmt.Println("middle")   //   |
	fmt.Println("nill")           //   |  direction of defer print
	defer fmt.Println("Exit")     //   |
	fmt.Println("Real Exit")      //   |
	defer fmt.Println("Shutdown") //   |
	fmt.Println("--------------")
	/*------------------
		Example 2:
	--------------------*/
	/*	res, err := http.Get("http://www.google.com/robots.txt")
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()
		robots, err := ioutil.ReadAll(res.Body)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", robots)
		fmt.Println("--------------")*/
	/*------------------
		Example 3:
	--------------------*/

	/*a := "start"
	defer fmt.Println(a)
	a = "end"*/

}
