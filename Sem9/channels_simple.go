package main

import (
	"fmt"
	"log"
)

func printChannelData(c chan int) {
	a := <- c
	//time.Sleep(3 * time.Second)
	fmt.Println("Data in channel is: ", a)
	log.Println("Finish")
 }

func main() {
	fmt.Println("Main started...")
	
	c := make(chan int) //Creo el canal
	
	// call to goroutine
	go printChannelData(c) //El proceso se espera hasta recibir la data 
	
	c <- 10 // Envia la data por el canal


	fmt.Println("Main ended...")
 }