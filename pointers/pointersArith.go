package main

import "fmt"

func main() {

	/*------------------
		Example 1:
	--------------------*/
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]

	fmt.Printf("%v %p %p\n", a, b, c)

	/* arithmetic is not allowed in go
	in other lenguages its posible in this
	case you can`t do &a[0] -4
	*/

	/*------------------
		Example 2:
	--------------------*/

	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	/*------------------
		Example 3:
	--------------------*/

	var ms2 *myStruct2
	fmt.Println(ms2) // here we have nill, in this case we inizialize the pointer to nil
	ms2 = new(myStruct2)
	fmt.Println(ms2) // here we initialize the var, the default value in Go is 0

	/*------------------
		Example 4:
	--------------------*/

	//the example bellow is a little creepy, in the next example,
	// we show a more readable code
	var ms3 *myStruct3
	ms3 = new(myStruct3)
	(*ms3).foo3 = 42
	fmt.Println((*ms3).foo3)

	// Bellow we have a better & readable code

	var ms4 *myStruct4
	ms4 = new(myStruct4)
	ms4.foo4 = 42
	fmt.Println(ms4.foo4)

	/*------------------
		Example 5:
	--------------------*/

	d := [3]int{1, 2, 3}
	e := d
	fmt.Println(d, e)

	f := []int{1, 2, 3} //example with slices
	g := f
	fmt.Println(f, g)
	f[1] = 42
	fmt.Println(f, g)

	/*------------------
		Example 6:
	--------------------*/

	h := map[string]string{"foo5": "bar", "baz": "buz"}
	i := h
	fmt.Println(h, i)
	h["foo5"] = "qux"
	fmt.Println(h, i)

}

type myStruct struct {
	foo int
}

type myStruct2 struct {
	foo2 int
}

type myStruct3 struct {
	foo3 int
}

type myStruct4 struct {
	foo4 int
}
