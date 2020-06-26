package utils

import "testing"

func TestStringInSlice(t *testing.T) {
	utils := Utils{}

	if !utils.stringInSlice("test", []string{"test"}) {
		t.Fatal("String must contains in array")
	}
}
