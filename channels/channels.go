package channels

import (
	"fmt"
	"math/rand"
	"time"
)



func chanTest() int{
	c := make(chan int)

	populate := func(i int){
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		// c <- i
	}
	for i := 0; i < 3; i++ {
		go populate(i)
	}
	fmt.Println(<-c)
	return <-c
}

func RunChannels(){
chanTest()
}

