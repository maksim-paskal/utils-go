package main

import "testing"

func TestStringInSlice(t *testing.T) {

	if StringInSlice("test", []string{"test"}) == false {
		t.Fatal("String must contains in array")
	}
}
