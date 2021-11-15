package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var remoteHost string
var localHost string
var n, min int // n tiene que sincronizarse

var chCont chan int // control de sincronizacion

func main() {
	//ingreso de puerto local
	bufferIn := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese el puerto del host local:")
	puerto, _ := bufferIn.ReadString('\n') // lee hasta presionar enter
	puerto = strings.TrimSpace(puerto) // Limpia de espacio al inicio y final
	localHost = fmt.Sprintf("localhost:%s",puerto) // localhost:8080

	//establecer direccion del host remoto
	fmt.Print("Ingrese el puerto del host remoto:")
	puerto, _ = bufferIn.ReadString('\n') // lee hasta presionar enter
	puerto = strings.TrimSpace(puerto) // Limpia de espacio al inicio y final
	remoteHost = fmt.Sprintf("localhost:%s",puerto) // localhost:8080

	//establecer numero a recibir ordenar
	fmt.Print("Ingrese el valor de n:")
	dato, _ := bufferIn.ReadString('\n') // lee hasta presionar enter
	dato = strings.TrimSpace(dato) // Limpia de espacio al inicio y final
	n, _ = strconv.Atoi(dato)  
	
	chCont = make(chan int, 1)
	chCont <- 0


	//Ejecucion
	ln, _ := net.Listen("tcp",localHost)
	defer ln.Close()

	//Manejo concurrente de conexiones
	for {
		con,_ := ln.Accept()
		go manejadorDeConexion(con)
		
	}

}

func manejadorDeConexion(con net.Conn){
	defer con.Close()

	bufferIn := bufio.NewReader(con)
	msg, _ := bufferIn.ReadString('\n')
	msg = strings.TrimSpace(msg)
	num, _ := strconv.Atoi(msg)

	fmt.Printf("Llegó el numero: %d\n",num)

	cont := <- chCont

	if cont == 0 {
		min = num		
	} else if num < min {
		enviar(min)
		min = num
	} else {
		enviar(num)
	}

	cont++
	//Evaluar
	if cont==n {
		fmt.Printf("El numero final es: %d\n",min)
		cont = 0
	} 

	chCont <- cont
}

func enviar(num int){
	con,_ := net.Dial("tcp",remoteHost)
	defer con.Close()
	fmt.Fprintln(con,num)

}