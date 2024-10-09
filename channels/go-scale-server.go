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
	Web1 = FakeSearch("web1", "The Go Programming language", "Https://golang.com")
	Web2 = FakeSearch("web2", "The Go Programming language", "Https://golang.com")
	Image1 = FakeSearch("image1", "The gopher", "Https://golang.png")
	Image2 = FakeSearch("image2", "The gopher", "Https://golang.png")
	Youtube1 = FakeSearch("video1", "Introduction to Golang", "Https://wwww.youtube.com/watch?v=some-video")
	Youtube2 = FakeSearch("video2", "Introduction to Golang", "Https://wwww.youtube.com/watch?v=some-video")
)

var (
	replicatedWeb = First(Web1, Web2)
	replicatedVideo = First(Youtube1, Youtube2)
	replicatedImage = First(Image1, Image2)
)

func First(replicas ...SearchFunc) SearchFunc{
	return func(query string) Result {
		c := make(chan Result, len(replicas));

		 sub_search := func(i int){
          c <- replicas[i](query)
		 }

		for i := range replicas {
          go sub_search(i)
		}	
	//  return the first results from the channel and discard the rest even if they are not ready
    return <-c
	}
}



func Search(query string, timeout time.Duration) ([]Result, error){
	resultschan := make (chan Result)
	timer := time.After(timeout)
    go func(){ resultschan <- Web1("Golang")}()
	go func(){ resultschan <- Youtube1("Golang")}()
	go func(){ resultschan <- Image1("Golang")}()

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


func SearchReplicated(query string, timeout time.Duration) ([]Result, error){
	resultschan := make (chan Result)
	timer := time.After(timeout)
    go func(){ resultschan <- replicatedWeb(query)}()
	go func(){ resultschan <- replicatedVideo(query)}()
	go func(){ resultschan <- replicatedImage(query)}()

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

    resp, err := SearchReplicated("golang", 80*time.Millisecond)
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