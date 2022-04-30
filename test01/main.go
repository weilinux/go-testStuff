package main

import "fmt"

func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func main() {
	fmt.Println("Go Testing demo")
	result := Calculate(8)
	fmt.Println(result)

}
