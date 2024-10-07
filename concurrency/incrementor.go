package concurrency

import (
	"fmt"
	"sync"
)


func RunIncrementor(){
 counter := 0
 var wg sync.WaitGroup
 var mu sync.Mutex

  wg.Add(100)
  for i := 0; i<100; i++ {
	go func(){
		mu.Lock()
      local := counter
	  local++
	  counter = local
	  fmt.Printf("Goroutine: %d, Counter: %d\n", i, counter)
	  mu.Unlock()
	  wg.Done()
	}()
  }
  wg.Wait()
  fmt.Println("End value: ", counter)
}



