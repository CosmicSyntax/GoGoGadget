package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var port string = "8000"

func main() {

	// Declaring the router...
	r := newRouter()

	defer http.ListenAndServe(":"+port, r)

	fmt.Println("Listening to port", port)

}

// newRouter is a constructor function that returns a router
// This allows for testing to see if router is working...
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", mainPage).Methods("GET")
	return r
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Hello World!")
}
