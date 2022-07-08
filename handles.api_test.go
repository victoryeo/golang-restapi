// handlers.api_test.go

package main

import (
	"net/http"
  "net/http/httptest"
  "testing"
	"io/ioutil"
	"strings"
)

// Test that a GET request to the home page returns the home page with
// the HTTP code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

  r.GET("/", getRoot)

  // Create a request to send to the above route
  req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
    // Test that the http status code is 200
    statusOK := w.Code == http.StatusOK
	
		// Test that the page title is "Home Page"
    p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}