package concurrency

import "fmt"


type human interface {
    speak()
}

type person struct {
  name string
}

func (p *person) speak(){
  fmt.Printf("%T: %v \n", p, p.name)
}

func saySomething(h human){
    h.speak()
}

func RunMethodSets(){
    p := person{
        name: "Andrew",
    }

    p.speak()
}