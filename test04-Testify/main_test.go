package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterUnique(t *testing.T) {
	input := []Developer{
		{Name: "Elliot"},
		{Name: "Elliot"},
		{Name: "David"},
		{Name: "Alex"},
		{Name: "Eva"},
		{Name: "Alan"},
	}

	expected := []string{
		"Elliot",
		"David",
		"Alex",
		"Eva",
		"Alan",
	}

	//1 testify way (simplify the testing process)
	assert.Equal(t, expected, FilterUnique(input))

	/*  //2 without testify
	result := FilterUnique(input)
	if ! reflect.DeepEqual(result, expected) {
		t.Fail()
	}
	*/
}

//negative test
func TestNegativeFilterUnique(t *testing.T) {
	input := []Developer{
		{Name: "Elliot"},
		{Name: "Elliot"},
		{Name: "David"},
		{Name: "Alex"},
		{Name: "Eva"},
		{Name: "Alan"},
	}

	expected := []string{
		"Elliot",
		"David",
		"Alan",
	}

	//there are many assert methods to simplify test
	assert.NotEqual(t, expected, FilterUnique(input))
}
