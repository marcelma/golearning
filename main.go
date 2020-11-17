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

	a := [4]int{1, 2, 3, 4}

	fmt.Printf("%v\t %T\n", a, a)

	x := [5]int{1, 2, 3, 4, 5}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)

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

	ex := []int{2, 3, 4, 5, 6, 7, 8, 9, 0, 1}

	fmt.Println(ex[:5])
	fmt.Println(ex[5:])
	fmt.Println(ex[2:7])
	fmt.Println(ex[1:6])

	m := []int{2, 3, 4, 5, 6, 7, 8, 9, 0, 1}
	m = append(m[:3], m[6:]...)
	fmt.Println(m)

	soa := make([]string, 50, 50)
	fmt.Println(soa)
	fmt.Println(len(soa))
	fmt.Println(cap(soa))

	soa = append(soa, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	for i := 0; i < len(soa); i++ {
		fmt.Println(i, soa[i])
	}

	h71 := []string{"James", "Bond", "Shaken, not stirred"}
	h72 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(h71)
	fmt.Println(h72)

	h7 := [][]string{h71, h72}
	for i, xs := range h7 {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v\n", j, val)

		}
	}
	fmt.Println(h7)

}

func mapExample() {
	headerFunc("mapExample")

	name := "Silva"

	m := map[string]int{
		"Marcell": 1,
		"Martini": 2,
	}

	h8 := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	// fmt.Println(h8)

	h8[`Marcell`] = []string{`Skydive`, `Swoop`, `4way`}

	delete(h8, `no_dr`)

	for k, names := range h8 {
		fmt.Println(k)
		for _, name := range names {
			fmt.Printf("\t%v\n", name)
		}
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
