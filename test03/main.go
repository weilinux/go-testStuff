package main

import "math"

//yet another way
type TestCase struct {
	value int
	expected bool
	actual bool
}
func CalculateIsArmstrongNumber(x int) bool {
	a := x /100
	b := x % 100 / 10
	c := x % 10
	return x == int(math.Pow(float64(a), 3)+math.Pow(float64(b),3)+math.Pow(float64(c), 3))
}
