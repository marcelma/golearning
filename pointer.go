package main

import (
	"fmt"
)

func pointerExamples() {
	headerOutPut("pointerExamples")

	a := 1
	// Endereço de a
	fmt.Println(&a)

	marcell := person{
		nome:      "Marcell",
		sobrenome: "Martini",
		idade:     41,
	}
	fmt.Println(marcell)

	changeMe(&marcell)
	fmt.Println(marcell)

}

func changeMe(p *person) {
	(*p).idade++
}
