package concurrency

import (
	"fmt"
	"time"
)



type Item struct {
	Title, Channel, GUID string
}

type Fetcher interface {
	Fetch() (items []Item, next time.Time, err error)
}

func (i Item) Fetch() (items []Item, next time.Time, err error){
  return nil, time.Now(), nil
}



func Fetch(uri string) Fetcher {
	i := new(Item)
	return *i
}


func RunAdvancedConcurrency(){
fmt.Println(time.Now())
}



