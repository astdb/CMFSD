package main

import (
	"fmt"
)

func main() {
	testStrings := []string{"", " ", `\0`, "This is a string", "\n"}

	for _, str := range testStrings {
		fmt.Print("IsNullOrEmpty(" + str + ") = ")

		if IsNullOrEmpty(str) {
			fmt.Println("TRUE")
		} else {
			fmt.Println("FALSE")
		}
	}
}

// IsNullOrEmpty takes a single string argument and returns a boolean value - true if the argument string is null or empty,
// true if not.
func IsNullOrEmpty(str string) bool {
	if str == "" {
		return true
	}

	if len(str) <= 0 {
		return true
	}

	return false
}
