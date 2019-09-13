package main

import "fmt"

type sizes struct {
	width  float32
	height float32
}

func main() {

	width := float32(0.0)
	height := float32(0.0)
	fmt.Println("Ingrese el alto")
	fmt.Scanf("%f", &width)
	fmt.Println("Ingrese el ancho")
	fmt.Scanf("%f", &height)
	sizes := newSizes(height, width)
	calcPerimeter(sizes)
	calcArea(sizes)
}

func newSizes(height float32, width float32) *sizes {
	p := sizes{height: height, width: width}
	return &p
}

func calcPerimeter(sizes *sizes) {
	fmt.Printf("El perimetro es %.2f", sizes.height*2+sizes.width*2)
}

func calcArea(sizes *sizes) {
	fmt.Printf("El area es %.2f", sizes.height*sizes.width)
}
