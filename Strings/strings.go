package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = []rune("résumé")
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("\nThe length of 'myString' is %v", len(myString))

	var myRune = 'a'
	fmt.Printf("\nThe value of 'myRune' is %v", myRune)

	var strSlice = []string{"s", "u", "b", "s", "t", "r", "i", "n", "g"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	fmt.Printf("\n%v", strBuilder.String())
	// cannot modify a string once created

}
