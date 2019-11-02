package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	encodedText := flag.String("c", "", "The encrypted text")

	flag.Parse()

	var subencoded = *encodedText

	if subencoded == "" {
		fmt.Println("Please run 'atbash' -h for usage information")
	} else {
		fmt.Println(atbash(subencoded))
	}

}

func atbash(s string) string {
	s = strings.ToLower(s)
	res := ""

	for _, c := range s {
		var decodedChar rune
		if c < rune('a') || c > rune('z') {
			decodedChar = c
		} else {
			diff := c - 'a'
			decodedChar = 'z' - diff
		}
		res = fmt.Sprintf("%s%c", res, decodedChar)
	}
	return res
}
