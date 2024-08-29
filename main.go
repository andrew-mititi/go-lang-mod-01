package main

import "fmt"

var service_name string  = "main.go"
const APP_NAME string = "golang module 1"

func ServiceName(name string) string {
	if name != "" {return name}
	return service_name
}

func main(){ 
	fmt.Printf("Running project %s\n", APP_NAME )
	fmt.Println("Hello from Go Mod one")
}