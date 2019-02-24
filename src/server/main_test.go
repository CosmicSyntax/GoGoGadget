package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// Form a new HTTP request
	req, err := http.NewRequest("GET", "", nil)

	// If we fail, stop the test
	if err != nil {
		t.Fatal(err)
	}

	// We use Go's httptest library to create an http recorder. This recorder
	// will act as the target of our http request... like a mini-browser
	recorder := httptest.NewRecorder()

	// Create an HTTP handler from mainPage function
	hf := http.HandlerFunc(mainPage)

	// Serve the HTTP request to our recorder
	hf.ServeHTTP(recorder, req)

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `Hello World!`
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", recorder.Body.String(), expected)
	}

}

func TestRouter(t *testing.T) {
	// Instantiate the router
	r := newRouter() // this is from the main.go

	// a mock server... this exposes the URL and making routing test possible
	fakeServer := httptest.NewServer(r)

	// make a get request... to the appropriate router
	res, err := http.Get(fakeServer.URL + "/hello")

	if err != nil {
		t.Fatal(err)
		// this will stop the test if error
	}

	// Ensure the get results in a status 200
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status should be 405, got %d", res.StatusCode)
	}

	// Test for the body
	byteD, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	expected := "Hello World!"

	if expected != string(byteD) {
		t.Errorf("Response should be %s, got %s", expected, string(byteD))
	}

}

func test404(t *testing.T) {
	r := newRouter()
	fakeServer := httptest.NewServer(r)
	res, err := http.Post(fakeServer.URL+"/test", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Ensure 404 works...
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status code should 404, but got %d", res.StatusCode)
	}

	// Try to read the body... should be empty
	byteD, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	expected := ""

	if expected != string(byteD) {
		t.Errorf("Response should be %s, got %s", expected, string(byteD))
	}

}
