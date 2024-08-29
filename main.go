package main

import (
	"fmt"

	"github.com/andrew-mititi/go-lang-mod-01/utils"
)

const APP_NAME string = "golang module 1"



func main(){ 
	fmt.Printf("Running project %s\n", APP_NAME )
	fmt.Printf("Hello from %s\n", utils.ServiceName("Module 1"))
}