package main

import (
	"fmt"
	"strings"
)

var service_name string  = "main.go"
const APP_NAME string = "golang module 1"

func ServiceName(name string) string {
	if name != "" {return strings.ToUpper(name)}
	return service_name
}

func main(){ 
	fmt.Printf("Running project %s\n", APP_NAME )
	fmt.Printf("Hello from %s\n", ServiceName("Module 1"))
}