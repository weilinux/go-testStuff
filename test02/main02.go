package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(Converter("this is some####sample string!@#4 you"))

}

func Converter(input string) (string, error) {
	if len(input) == 7 {
		return "", fmt.Errorf("input must be not length 7 characters")
	}
	splitter := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	splits := strings.FieldsFunc(input, splitter)
	out := strings.Join(splits, "_")

	return out, nil
}
