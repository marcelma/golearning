package main

import (
	"fmt"
)

func main() {
	arrayExample()
	sliceExample()
	mapExample()
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

func mapExample() {
	headerFunc("mapExample")

	name := "Silva"

	m := map[string]int{
		"Marcell": 1,
		"Martini": 2,
	}

	// comma ok idiom
	v, ok := m[name]
	fmt.Println(v)
	fmt.Println(ok)

	// comma ok idiom
	if v, ok := m[name]; ok {
		fmt.Println(v)
	}

	// add element to map
	m["Skydive"] = 3
	fmt.Println(m)

	// delete element from map
	delete(m, "Marcell")
	fmt.Println(m)
	delete(m, "Silva")

	if i, ok := m["Silva"]; ok {
		fmt.Println(i)
	}
}
