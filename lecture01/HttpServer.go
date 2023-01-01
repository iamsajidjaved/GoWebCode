package main

import (
	"fmt"
	"net/http"
)

func main() {
	// to know more about http package https://gobyexample.com/http-servers
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// to know more about fmt package https://www.geeksforgeeks.org/fmt-package-in-golang/
	fmt.Fprintf(w, "Welcome to go lang. My 1st app is servering on URL path %s", r.URL.Path)
}
