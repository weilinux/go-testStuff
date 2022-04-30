package main

import "testing"

func TestConverter(t *testing.T) {
	input := "this is a sample string"
	expectedOutput:= "this_is_a_sample_string"
	actual, err := Converter(input)
	if err != nil {
		t.Errorf("extpected no error, but got %v", err)
	}

	if actual != expectedOutput {
		t.Errorf("expected %v got %v", expectedOutput, actual)
	}

}
