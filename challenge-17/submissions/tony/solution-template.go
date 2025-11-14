package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// Get input from the user
	var input string
	fmt.Print("Enter a string to check if it's a palindrome: ")
	fmt.Scanln(&input)

	// Call the IsPalindrome function and print the result
	result := IsPalindrome(input)
	if result {
		fmt.Println("The string is a palindrome.")
	} else {
		fmt.Println("The string is not a palindrome.")
	}
}
// IsPalindrome checks if a string is a palindrome.
// A palindrome reads the same backward as forward, ignoring case, spaces, and punctuation.
func IsPalindrome(str string) bool {
	// TODO: Implement this function

	// 1. Clean the string (remove spaces, punctuation, and convert to lowercase)
	n := []rune(Normalize(str))
	// 2. Check if the cleaned string is the same forwards and backwards
	l:=len(n) - 1
	if l > 0 {
    	mid:= l/2
	    for i := 0; i<=mid;i++ {
		    if n[i] != n[l-i] {
			    return false
		    }			
	    }
	}
	return true
}
func Normalize(str string) string {
	s:=[]rune(str)
	var result strings.Builder
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			result.WriteRune(unicode.ToLower(c))
		}
	}
	return result.String()
}

