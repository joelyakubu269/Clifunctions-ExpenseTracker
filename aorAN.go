package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(aoran("A honest man n boy, a egg, A else"))
}
func aoran(str string) string {
	words := strings.Fields(str)
	for i := 0; i < len(words); i++ {
		if strings.ToLower(words[i]) == "an" && !strings.ContainsAny(words[i+1][:1], "aeiouhAEIOUH") {
			words[i] = "a"
		}
		if strings.ToLower(words[i]) == "a" && strings.ContainsAny(words[i+1][:1], "aeiouhAEIOUH") {
			words[i] = "an"
		}
	}
	return strings.Join(words, " ")
}
