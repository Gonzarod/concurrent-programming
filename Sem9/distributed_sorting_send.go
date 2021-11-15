package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var remotehost string

func main() {
	//modo cliente, envío
	buffIn := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese el puerto remoto: ")
	puerto, _ := buffIn.ReadString('\n')
	puerto = strings.TrimSpace(puerto)
	remotehost = fmt.Sprintf("localhost:%s", puerto)
	con, _ := net.Dial("tcp", remotehost) //dirección de red + puerto donde se va a comunicar
	defer con.Close()
	//envio de datos al server
	enviar(60)
	enviar(30)
	enviar(80)
	enviar(10)
}

func enviar(num int) {
	con, _ := net.Dial("tcp", remotehost)
	defer con.Close()
	fmt.Fprintln(con, num)
}
