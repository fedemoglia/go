package main

import (
	"fmt"
	"strings"
)

func main() {
	word := ""
	fmt.Println("Ingresa una palabra")
	fmt.Scanf("%s", &word)
	fmt.Println(change(word))
}

func change(word string) string {
	return strings.Replace(strings.ToLower(word), "o", "a", -1)
}
