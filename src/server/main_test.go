package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// Instantiate the router
	r := newRouter()

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
