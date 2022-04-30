package main

import "testing"

func TestFilterUnique(t *testing.T) {
	input := []Developer{
		Developer{Name: "Elliot"},
		Developer{Name: "Elliot"},
		Developer{Name: "David"},
		Developer{Name: "Alex"},
		Developer{Name: "Eva"},
		Developer{Name: "Alan"},
	}

	expected := []string{
		"Elliot",
		"david",
		"Alex",
		"Eva",
		"Alan",
	}

	result := FilterUnique(input)




}