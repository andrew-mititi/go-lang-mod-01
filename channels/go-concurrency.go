package channels

import (
	"fmt"
	"math/rand"
	"time"
)


func RunChannelsConcurrency(){
  ch := gennums(1,2,3,4)
  for n := range ch {   // Read values from the channel until it is closed
	fmt.Println(n)
}
}


func gennums(nums ...int) <-chan int {
   
	out := make(chan int)

	inter_gen := func(i int){
		defer wg.Done()
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		out <- i
	}

	for _, n := range nums {
		wg.Add(1)	
		go inter_gen(n)
	}

	go func() {
		wg.Wait()            // Wait for all goroutines to finish
		close(out)           // Close the channel after all sends are done
	}()
    fmt.Println("Concurrency: Gen function done running")
	return out
}
