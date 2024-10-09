package main

import "github.com/andrew-mititi/go-lang-mod-01/channels"

const APP_NAME string = "golang module 1"

type Person struct {
	firstname string
	age       int
}

type personMap map[string]*Person

var persons = personMap{
	"andrew": &Person{
		firstname: "Andrew",
		age: 26,
	},
}

// func (p Person) speak() {
	

// }

func main() {
	// ps := persons["andrew"]
	// fmt.Println(ps.firstname)
	// ps.firstname = "Mititi"
	// fmt.Println(persons["andrew"].firstname)
	// concurrency.RunIncrementor()
	channels.RunChannels()
}

