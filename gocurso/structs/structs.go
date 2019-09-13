package structs

import "fmt"

// GetArray imprimie dos arrays
func GetArray() {
	var arr1 [2]string
	arr2 := [3]int{1, 2, 3}
	arr1[0] = "array"
	arr1[1] = "array2"
	fmt.Println(arr1, arr2)
}

// GetSlice imprime un slice
func GetSlice() {
	var slice1 []string
	slice1 = append(slice1, "mi", "slice", "1")
	fmt.Println(slice1)
}

// PlatziCourse es una estructura
type PlatziCourse struct {
	Name   string
	Slug   string
	Skills []string
}

// Subscribe es un método de PlatziCourse
func (p PlatziCourse) Subscribe(name string) {
	fmt.Printf("La persona %s se ha registrado al curso %s\n", name, p.Name)
}

// PlatziCareer es una estructura que toma los métodos de PlatziCourse
type PlatziCareer struct {
	PlatziCourse
}

// Subscribe registra una persona a una carrera
func (p PlatziCareer) Subscribe(name string) {
	fmt.Printf("La persona %s se ha registrado a la carrera %s\n", name, p.Name)
}
