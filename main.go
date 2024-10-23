package main

import "github.com/andrew-mititi/go-lang-mod-01/concurrency"

const APP_NAME string = "golang module 1"

func main() {
	// http.HandleFunc("/hello", channels.HandleHello)
	// http.HandleFunc("/search", channels.HandleSearch)

	// fmt.Println("Server listening on port 9000")
	// log.Fatal(http.ListenAndServe("localhost:9000", nil))
  concurrency.RunAdvancedConcurrency()
}