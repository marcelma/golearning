package main

import (
	"fmt"
)

func main() {
	arrayExample()
	sliceExample()
}

func headerFunc(h string) {
	fmt.Printf("\n######################\n")
	fmt.Printf("# %v\n", h)
	fmt.Printf("######################\n")
}

func arrayExample() {

	headerFunc("arrayExample")

	a := []int{1, 2, 3, 4}

	fmt.Printf("%v\t %T\n", a, a)
}

func sliceExample() {
	headerFunc("sliceExample")

	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)

	z := append(x[0:2], x[4:]...)
	fmt.Println(z)
}
