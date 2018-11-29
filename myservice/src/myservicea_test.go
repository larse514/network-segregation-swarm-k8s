package main

import "testing"

const (
	code    = "200"
	message = "SOMEMESSAGE"
)

//Dummy test to demonstrate test in ci pipeline
func TestMyResponse(t *testing.T) {
	response := MyResponse{Code: code, Message: message}

	if response.Code != code {
		t.Log("expected code ", code, " but found ", response.Code)
		t.Fail()
	}
}
