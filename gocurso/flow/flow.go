package flow

import "fmt"

// IfTest verifica si un número es par o impar
func IfTest() {
	var number = 0
	fmt.Print("Ingresa un número del 1 al 10: ")
	fmt.Scanf("%d", &number)
	if number%2 == 0 {
		fmt.Println("El número es par")
	} else {
		fmt.Println("El número es impar")
	}

	if number2 := 3; number2 == 3 {
		fmt.Println("Entró al condicional")
	}
}

// SwitchTest solicita un número del 1 al 10 y lo evalúa
func SwitchTest() {
	var number = 0
	fmt.Print("[Switch] Ingresa un número del 1 al 10: ")
	fmt.Scanf("%d", &number)

	switch number {
	case 1:
		fmt.Println("El número es 1")
	default:
		fmt.Println("El número no es 1")
	}

	switch {
	case IsPar(number):
		fmt.Println("El número es par")
	default:
		fmt.Println("El n")
	}
}

func IsPar(number int) bool{
	return number%2==0
}

