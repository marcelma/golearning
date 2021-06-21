package main

import (
	"fmt"

	"github.com/marcelma/golearning/basic/testing/examples/quote"
	"github.com/marcelma/golearning/basic/testing/examples/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
