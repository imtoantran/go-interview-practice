package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read input from standard input
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		// Call the ReverseString function
		output := ReverseString(input)

		// Print the result
		fmt.Println(output)
	}
}

// ReverseString returns the reversed string of s.
func ReverseString(s string) string {
	// TODO: Implement the function
	str:=[]rune(s)
	str_len := len(str) - 1
	var revert strings.Builder
	for i:= str_len;i>=0; i-- {
	    revert.WriteRune(str[i])
	}
	return revert.String()
}
