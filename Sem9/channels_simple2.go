package main

import (
	"fmt"
	"time"
)

func enviarNumero(ch chan int) {
	for i := 0 ; i < 10 ; i++ {  // Envio datos limitados, por eso necesito close()
		ch <- i
		time.Sleep(3*time.Second)
	}
	close(ch)
 }

func main() {	
	ch := make(chan int)

	go enviarNumero(ch)

	for num := range ch { //Lee hasta que se termine de enviar los datos
		fmt.Println(num)
	}

 }