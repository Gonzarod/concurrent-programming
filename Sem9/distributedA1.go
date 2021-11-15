package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// Nodos: maquina o servidor donde esta corriendo un programa

//Servidor esta escuchando peticiones. Tiene un ip y escucha por un puerto
func main(){
	//rol SERVIDOR------------------------------------------------------- A
	ln, err := net.Listen("tcp","localhost:9100")

	if err != nil {
		fmt.Println("Fallo algo", err.Error())
		os.Exit(1) // Finaliza programa
	}

	defer ln.Close()
	
	con, _ := ln.Accept()

	defer con.Close() 
	//leer lo que   envia el cliente
	bufferIn := bufio.NewReader(con)

	msg, _ := bufferIn.ReadString('\n')

	fmt.Println(msg)
	

	//Respondemos -- al cleinte
	con2, err := net.Dial("tcp","localhost:9200")

	defer con2.Close()

	fmt.Fprintln(con2,"recibido !!!!!")


}