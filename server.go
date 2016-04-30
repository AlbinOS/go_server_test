package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	t := time.Now()
	fmt.Println(t)
	fmt.Println("Wooty woot !")
	fmt.Println("------------")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4030", nil)
}
