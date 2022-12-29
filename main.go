package main

import (
	"fmt"
	"net/http"
	"railway/config"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {	
	config.LoadConfig()
	p := fmt.Sprintf(":%s", config.APIPort)
	http.HandleFunc("/", greet)
	http.ListenAndServe(p, nil)
}
