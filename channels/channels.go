package channels

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str string
	wait chan bool
}

func boring(msg chan Message){
    waitForit := make(chan bool)
	go func(){
		for i := 0; ; i++ {
		   msg <- Message{ str: fmt.Sprintf("%s, %d", msg, i), wait: waitForit }
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForit
		}
	}()
	
}

func fanIn(input1, input2 <-chan string) <-chan string{
	c := make(chan string)
	go func(){for { c <- <-input1}}()
	go func(){for { c <- <-input2}}()
	return c
}


func RunChannels(){

}

