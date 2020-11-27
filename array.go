package main

import "fmt"

func arrayExample() {

	headerOutPut("arrayExample")

	a := [4]int{1, 2, 3, 4}

	fmt.Printf("%v\t %T\n", a, a)

	x := [5]int{1, 2, 3, 4, 5}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)

}
