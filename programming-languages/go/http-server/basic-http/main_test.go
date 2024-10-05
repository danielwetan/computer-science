package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("helloHandler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello World"
	if rr.Body.String() != expected {
		t.Errorf("helloHandler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestTestHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("testHandler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Test"
	if rr.Body.String() != expected {
		t.Errorf("testHandler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
