package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Lectura de Datos
	//

	//bloque textual multilinea
	menu:=`
			Bienvenido
			[1] Pizza
			[2] Empanada
			¿Que prefieres?
	`

	fmt.Println(menu)

	bufferIn:= bufio.NewReader(os.Stdin)
	datoIn, _:= bufferIn.ReadString('\n')
	dato:= strings.Trim(datoIn,"\r\n")

	switch dato {
	case "1":
		fmt.Println("Elegiste Pizza")
	case "2":
		fmt.Println("Elegiste Empanada")
	default:
		fmt.Println("Opción inválida!!!")
	}

}
