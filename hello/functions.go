package main

import "fmt"

const helloWorld string = "Hola %s %s, bienvenido al fascinante mundo de Go.\n"
const enterANumber = "Ingresa un número: "

func main() {
	lastname := "<Modificar con mi apellido>"
	name := getName()
	number := sum(50, 50)
	a, b, c := getVariables()
	fmt.Printf(helloWorld, name, lastname)
	fmt.Print("Hola mundo ")
	fmt.Println(number, a, b, c)
	operando1, operando2 := getOperandos()
	fmt.Println(operando1, operando2)
}

func getName() string {
	var name string
	name = "Sin nombre"
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	return name
}

func getVariables() (int, string, int) {
	return 1, "Fede", 3
}

func sum(a int, b int) int {
	return a + b
}

func getOperandos() (int, int) {
	number1 := 0
	number2 := 0
	fmt.Print(enterANumber)
	fmt.Scanf("%d", &number1)
	fmt.Print(enterANumber)
	fmt.Scanf("%d", &number2)
	return number1, number2
}
