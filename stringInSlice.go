package main

import "testing"

func TestStringInSlice(t *testing.T) {

	if !StringInSlice("test", []string{"test"}) {
		t.Fatal("String must contains in array")
	}
}
