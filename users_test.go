package main

import (
	"bytes"
	"testing"
	"net/http"
	"net/http/httptest"
	"build-rest-api/controllers"
	"fmt"
)

func TestRegisterUser(t *testing.T) {
	var requestBody = []byte(`{"email": "xxx1@xx.com", "password": "12121212"}`)
	req, err := http.NewRequest("POST", "/api/user/new", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.CreateAccount)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}


}
