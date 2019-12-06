package main

import (
	"testing"
)

func TestValidation(t *testing.T) {
	//test case to check pass
	test := "379354508162306"

	if !Validation(test) {
		t.Errorf("Expected ”%s” to be valid according to the Luhn algorithm", test)
	}

	test = "4388576018402626"

	if Validation(test) {
		t.Errorf("Expected ”%s” to be invalid according to the Luhn algorithm", test)
	}	
	
	//test case to check fail
	
	test = ""

	if Validation(test) {
		t.Errorf("Expected ”%s” to be invalid according to the Luhn algorithm", test)
	}
	
	test = "4388576018400000000000000"

	if Validation(test) {
		t.Errorf("Expected ”%s” to be invalid according to the Luhn algorithm", test)
	}
}
