package main

import (
	"testing"
)

func TestValid(t *testing.T) {
	test := "1111111116"

	if !Valid(test) {
		t.Errorf("Expected ”%s” to be valid according to the Luhn algorithm", test)
	}

	test = "1111111111"

	if Valid(test) {
		t.Errorf("Expected ”%s” to be invalid according to the Luhn algorithm", test)
	}
}
