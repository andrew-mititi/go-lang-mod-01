package concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func RunRoutines() {
	fmt.Println("Starting routines ...")
    wg.Add(2)
	fmt.Println(runtime.NumGoroutine())

	go foo()
	go bar()
	fmt.Println(runtime.NumGoroutine())

	fmt.Println("Done")
	wg.Wait()

}

func bar() {
	fmt.Println("Bar")
	wg.Done()
}

func foo() {
	fmt.Println("Foo")
	wg.Done()
}
