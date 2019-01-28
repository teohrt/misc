package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

/*
These tests can be run in parallel,
since they depend on the same
server go routine
*/

func TestGetBooks(t *testing.T) {
	go initServer()
	t.Parallel()
	// This go routine is now available for other tests.

	url := "http://localhost:8000/api/books"
	resBody, statusCode, err := httpRequestHelper("GET", url, nil)
	if err != nil {
		t.Error(err)
		return
	}
	if statusCode != 200 || len(resBody) == 0 {
		t.Errorf("GET request errored. Status Code: %d\nResponse Body: \n%s\n", statusCode, resBody)
		return
	}
	log.Printf("PASSED: TestGetBooks Status Code: %d\nResponse Body: \n%s\n", statusCode, resBody)
	return
}

func TestGetBook(t *testing.T) {
	idToRetrieve := "1"
	url := "http://localhost:8000/api/books/" + idToRetrieve

	resBody, statusCode, err := httpRequestHelper("GET", url, nil)
	if err != nil {
		t.Error(err)
		return
	}
	if statusCode != 200 || len(resBody) == 0 {
		t.Errorf("GET request errored. Status Code: %d\nResponse Body: \n%s\n", statusCode, resBody)
		return
	}
	log.Printf("PASSED: TestGetBook Status Code: %d\nResponse Body: \n%s\n", statusCode, resBody)
	return
}

func TestCreateBook(t *testing.T) {
	url := "http://localhost:8000/api/books"
	var newBook = []byte(`{"isbn": "123456", "title": "This is a Test", "author": {"firstname": "Billy-Bob", "lastname": "Joe"}}`)

	previosLength := len(books)
	resBody, statusCode, err := httpRequestHelper("POST", url, newBook)
	if len(books) != previosLength+1 || err != nil || statusCode != 200 {
		t.Errorf("POST request failed. Status Code: %d", statusCode)
		return
	}
	log.Printf("TestCreateBook Status Code: %d\nResponse Body: \n%s\n", statusCode, resBody)
	return
}

func TestDeleteBook(t *testing.T) {
	idToDelete := "1"
	url := "http://localhost:8000/api/books/" + idToDelete

	resBody, statusCode, err := httpRequestHelper("DELETE", url, nil)
	if err != nil {
		t.Error(err)
		return
	}
	if statusCode != 200 || len(resBody) == 0 {
		t.Errorf("DELETE request errored. Status Code: %d\nResponse Body: \n%s\n", statusCode, resBody)
		return
	}
	log.Printf("PASSED: TestDeleteBook Status Code: %d\nResponse Body: \n%s\n", statusCode, resBody)
	return
}

func TestUpdateBook(t *testing.T) {
	idToUpdate := "2"
	var updatedBook = []byte(`{"isbn": "10110100", "title": "This book was updated", "author": {"firstname": "People", "lastname": "Change"}}`)
	url := "http://localhost:8000/api/books/" + idToUpdate

	resBody, statusCode, err := httpRequestHelper("PUT", url, updatedBook)
	if err != nil {
		t.Error(err)
		return
	}
	if statusCode != 200 || len(resBody) == 0 {
		t.Errorf("PUT request errored. Status Code: %d\nResponse Body: \n%s\n", statusCode, resBody)
		return
	}
	log.Printf("PASSED: TestUpdateBook Status Code: %d\nResponse Body: \n%s\n", statusCode, resBody)
	return
}

/*
Helper function for HTTP requests
Returns response body, status code, and error values
(Written specifically for PUT and DELETE request functionality.)

reqType paramater can be: "GET", "POST", "PUT", "DELETE"
body parameter can be nil
*/
func httpRequestHelper(reqType string, url string, body []byte) (string, int, error) {

	client := &http.Client{}
	bodyBuffer := bytes.NewBuffer(body)

	req, err := http.NewRequest(reqType, url, bodyBuffer)
	if err != nil {
		return "", -1, err
	}
	res, err := client.Do(req)
	if err != nil {
		return "", -1, err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", res.StatusCode, err
	}
	return string(resBody), res.StatusCode, err
}
