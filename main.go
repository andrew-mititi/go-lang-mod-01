package main

import (
	"fmt"

	utils_helpers "github.com/andrew-mititi/go-lang-mod-01/utils-helpers"
)

const APP_NAME string = "golang module 1"



func main(){ 
	fmt.Printf("Running project %s\n", APP_NAME )
	fmt.Printf("Hello from %s\n", utils_helpers.ServiceName(APP_NAME))
}