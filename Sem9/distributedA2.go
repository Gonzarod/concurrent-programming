package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)


func main(){
	//rol CLIENTE ------------------------------------------------------------------ B
	con, _ := net.Dial("tcp","localhost:9100")
	defer con.Close()

	//Enviar mensaje al servidor
	fmt.Fprintln(con, "Saludos desde el cliente")
	

	// Escuchar lo que va a responder el Servidor -------------------------------------------------------------------------------
	ln, err := net.Listen("tcp","localhost:9200")

	if err != nil {
		fmt.Println("Fallo algo", err.Error())
		os.Exit(1) // Finaliza programa
	}

	defer ln.Close()
	
	con2, _ := ln.Accept()

	defer con.Close() 
	//leer lo que envia el servidor
	bufferIn := bufio.NewReader(con2)

	msg, _ := bufferIn.ReadString('\n')

	fmt.Println(msg)
}