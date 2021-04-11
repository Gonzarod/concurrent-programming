package main

import "fmt"

func main() {
	//
	var large float64=12.26
	//
	largeSimple:=12



	fmt.Println("Large Number es: ",large)
	fmt.Println("Simple Version: ",largeSimple)
	fmt.Println("Multiplicacion: ",large*float64(largeSimple))
	fmt.Println("El valor del largo es mayor al valor del ancho: ",large>float64(largeSimple))
}
