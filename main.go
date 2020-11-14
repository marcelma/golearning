package main

import (
	"fmt"
)

func main() {
	arrayExample()
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
