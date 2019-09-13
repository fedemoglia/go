package main


import (
	"fmt"
	"strings"
	"runtime"
)

func main() { 
	strings2()
	switchTest()
	getOS()
}

func strings2() {
	fmt.Println(strings.ToUpper("hola"))
	fmt.Println(strings.Count("hola", "a"))
}

func switchTest() {
	var number = 0
	fmt.Println("Ingrese un número")
	fmt.Scanf("%d", &number)

	switch  {
	case number%2 == 0: 
		fmt.Println("El número es par")
	case number%2 != 0:
		fmt.Println("El número es impar")
	}
}

func getOS() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}