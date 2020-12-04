package main

import (
	"fmt"
	"io"
	"os"
)

func stdoutExamples() {
	headerOutPut("pointerExamples")

	fmt.Println("Hello, playground")
	fmt.Fprintln(os.Stdout, "Hello, playground")
	io.WriteString(os.Stdout, "Hello, playground\n")
}
