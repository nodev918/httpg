package main

import (
	"fmt"
	"net/http"
)


func indexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("hi")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "hello world")

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/route",func(w http.ResponseWriter, r *http.Request){
		
	})
	http.ListenAndServe(":8888",nil)
}

