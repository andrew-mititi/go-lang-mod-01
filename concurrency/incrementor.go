package concurrency

import (
	"fmt"
	"runtime"
	"sync"
)


func RunIncrementor(){
 counter := 0
 var wg sync.WaitGroup

  wg.Add(100)
  for i := 0; i<100; i++ {
	go func(){
      local := counter
	  runtime.Gosched()
	  local++
	  counter = local
	  fmt.Printf("Goroutine: %d, Counter: %d\n", i, counter)
	  wg.Done()
	}()
  }
  wg.Wait()
  fmt.Println("End value: ", counter)
}