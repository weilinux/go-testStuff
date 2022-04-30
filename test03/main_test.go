package main

import "testing"


func TestCalculateIsBoolean(t *testing.T) {
	testCase := TestCase{
		value: 371,
		expected: true,
	}
	testCase.actual = CalculateIsArmstrongNumber(testCase.value)
	if testCase.actual != testCase.expected {
		t.Fail()
	}
}

func TestCalculateIsArmstrongNumber(t *testing.T) {

	//multiple t.Run case, or  put in for range
	t.Run("should return true for 371", func(t *testing.T) {
		testCase := TestCase{
			value: 371,
			expected: true,  //expected: false
		}

		testCase.actual = CalculateIsArmstrongNumber(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})

	t.Run("should return true for 370", func(t *testing.T) {
		testCase := TestCase{
			value: 370,
			expected: true, //expected false
		}

		testCase.actual = CalculateIsArmstrongNumber(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})
}
