package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

type Animal struct {
	Name string `requiered max: "100"`
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jonh Pertwee",
		companions: []string{
			"Liz shaw",
			"Jo Grant",
			"Sara Jane Smith",
		},
	}
	fmt.Println(aDoctor.number)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions)
	fmt.Println(aDoctor.companions[1])

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

}
