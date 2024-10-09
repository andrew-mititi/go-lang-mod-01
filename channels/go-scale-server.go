package channels

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)


type Result struct{
	Title, URL string
}
type Response struct {
	Elapsed time.Duration
	Results []Result
}


type SearchFunc func(query string) Result

var (
	Web = FakeSearch("web", "The Go Programming language", "Https://golang.com")
	Image = FakeSearch("image", "The gopher", "Https://golang.png")
	Youtube = FakeSearch("video", "Introduction to Golang", "Https://wwww.youtube.com/watch?v=some-video")
)

func First(replicas ...SearchFunc) SearchFunc{
	return func(query string) Result {
		c := make(chan Result, len(replicas));


		for i := range replicas {
         go func(){
			res := replicas[i](query)
			c <- res
		 }()
		}
    return <-c
	}
}

func Search(query string, timeout time.Duration) ([]Result, error){
	resultschan := make (chan Result)
	timer := time.After(timeout)
    go func(){ resultschan <- Web("Golang")}()
	go func(){ resultschan <- Youtube("Golang")}()
	go func(){ resultschan <- Image("Golang")}()

	var results []Result

	for i := 0; i < 3; i++ {
     select {
	 case result := <-resultschan:
	   results = append(results, result)
	 case <-timer:
		return results, errors.New("Timed out")

	 }
	} 
	return results, nil
}

func FakeSearch(kind, title, url string) SearchFunc{
	return func(query string) Result{
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{
			Title: fmt.Sprintf("%s(%q): %s", kind, query, title),
			URL: url,
		}
	}
}


func HandleHello(w http.ResponseWriter, req *http.Request){
	log.Println("serving", req.URL)
	fmt.Fprintln(w, "Hello World")
}

func HandleSearch(w http.ResponseWriter, req *http.Request){
	queryParam := req.FormValue("q")

	if(queryParam == ""){
		http.Error(w, `Missing "q" URL Parameter`, http.StatusBadRequest)
		return
	}

	start := time.Now()

    resp, err := Search("golang", 80*time.Millisecond)
   if(err != nil) {
	 http.Error(w, err.Error(), http.StatusInternalServerError)
   }
   elapsed := time.Since(start)

   response := Response{
     Elapsed: time.Duration(elapsed.Milliseconds()),
	 Results: resp,
   }
   log.Println(fmt.Printf("Elapsed time: %d ", elapsed.Milliseconds()))
	json.NewEncoder(w).Encode(response)
	

}