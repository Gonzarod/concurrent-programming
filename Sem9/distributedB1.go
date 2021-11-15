package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
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

	for {
		con, _ := ln.Accept()
		go manejador(con)
	}

	

}

func manejador(con net.Conn)  {

	defer con.Close()
	//Recuperar mensaje
	bufferIn := bufio.NewReader(con)
	msg, _ := bufferIn.ReadString('\n')

	fmt.Printf("Procesando petición de, %s", msg)
	time.Sleep(5 * time.Second)
	//Respondiendo al cliente por la misma conexion
	fmt.Fprintln(con, "Peticion Procesada "+msg)
		
}