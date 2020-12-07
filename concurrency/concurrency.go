package concurrency

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/marcelma/golearning/tools"
)

var wg sync.WaitGroup

func fooC() {
	for i := 0; i < 10; i++ {
		fmt.Println("fooC:", i)
	}
	wg.Done()
}

func barC() {
	for i := 0; i < 10; i++ {
		fmt.Println("barC:", i)
	}
}

func exemplo1() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("Arch\t\t", runtime.GOARCH)
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("goRoutine\t", runtime.NumGoroutine())

	wg.Add(1)
	go fooC()
	barC()

	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("goRoutine\t", runtime.NumGoroutine())
	wg.Wait()
}

// Examples of concurrency
func Examples() {
	tools.HeaderOutPut("concurrencyExample")

	exemplo1()
}
