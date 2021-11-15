package main

import "fmt"

func main() {
	//ch := make(chan int) // Canal Sincrono . Solo permite que los procesos o bien puedan leer o bien escribe . No puede hacer los dos a la vez
	ch := make(chan int, 1) // Canal Asincrono. Puedes escribir y leer a la vez. El parametro indica cuantas veces se puede hacer

	ch <- 5 // envio

	fmt.Printf("%d\n", <- ch) // leo

	//defer : Declaracion de funciones diferidas. Se utiliza bastante en los methodos de comunicación distribuidas
	// porque tenemos quee cerrar conexiones (garantizamos que siempre se cierre al final)

}