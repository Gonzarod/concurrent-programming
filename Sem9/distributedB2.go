package main

import (
	"bufio"
	"fmt"
	"net"

)


func main(){
	//rol CLIENTE ------------------------------------------------------------------ B
	
	//Enviar mensaje al servidor
	con, _ := net.Dial("tcp","localhost:9100")
	fmt.Fprintln(con, "Jesus")

	defer con.Close()

	//Leyendo respuesta de
	bufferIn := bufio.NewReader(con)
	msg, _ := bufferIn.ReadString('\n')
	fmt.Print(msg)
	
	
	
}