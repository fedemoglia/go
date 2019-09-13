package main

import (
	"fmt"
	/*"gocurso/flow"
	"gocurso/name"
	"gocurso/numbers"
	."gocurso/structs"
	"gocurso/strings2"**/
	"gocurso/maps"
	//"gocurso/flow"
)

const helloWorld string = "Hola %s %s, bienvenido al fascinante mundo de Go.\n"
const testConst = "Test"

func main() {
	defer fmt.Println("último que ejecuto")
	/*lastname := "<Modificar con mi apellido>"
	number := numbers.Sum(50, 50)
	firstName := name.GetName()
	a, b, c := numbers.GetVariables()
	f32, f64 := numbers.GetFloat()
	stringUTF8 := name.GetUnicode()
	fmt.Printf(helloWorld, firstName, lastname)
	fmt.Print("Hola mundo")
	fmt.Println(number, a, b, c)
	fmt.Println(f32, f64)
	fmt.Println("Cadena con UTF8:", stringUTF8)
	fmt.Println(string("hola"[1]))
	fmt.Println("Cantidad de letras que tiene hola ", len("hola"))
	GetArray()
	GetSlice()
	flow.IfTest()

	strings2.Strings2()
	flow.SwitchTest()*/
	fmt.Println(maps.GetMap())
}
