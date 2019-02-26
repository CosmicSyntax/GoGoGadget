package main

import (
	"fmt"
	"net/http"
	"time"

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
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/hello", mainPage).Methods("GET")

	r.HandleFunc("/name", name).Methods("GET")

	sF := http.Dir("/Users/choius.ibm.com/Documents/GoGoGadget/src/server/static")

	// below ensure when typing ./static/index.html ... it doesnt' translate to ./static/static/index.html
	sH := http.StripPrefix("/static", http.FileServer(sF))

	// tell router to use staticFileHandler for all routers for assets
	r.PathPrefix("/static").Handler(sH).Methods("GET")

	r.Handle("/", http.FileServer(sF))

	return r
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Hello World!")
}

func name(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, time.Now())
}
