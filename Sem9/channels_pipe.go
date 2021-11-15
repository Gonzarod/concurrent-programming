package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MAX = 9
	K   = 4
)

func compress(inC, pipe chan rune) {
	n := 0
	previous := <- inC

	for {
		// a partir del siguiente
		c := <- inC
		
		if c==previous &&  n < MAX - 1 {
			n++
		} else {
			if n > 0 {
				pipe <- rune(n + 49)
				n = 0
			}
			pipe <- previous
			previous = c
		}
	}
}

func output(pipe, outC chan rune) {
	m:= 0
	for {
		outC <- <- pipe
		m++
		if m >= K {
			outC <- '\n'
			m = 0 
		}  
	}
}

func main() {
	inC := make(chan rune)
	pipe := make(chan rune)
	outC := make(chan rune)

	go compress(inC, pipe)

	go output(pipe, outC)
	
	go func ()  { // Proceso independiente enfocado a enviar Datos por el canal inC
		m := 1
		rand.Seed(time.Now().UTC().UnixNano())	
		for {
			m++
			if m%5 == 0 {
				time.Sleep(4*time.Second)
			}
			inC <- rune(rand.Intn(26)+65) //65 a 90 
		}
	}()

	for {
		fmt.Printf("%c", <-outC) // espera a que llegue algo por el canal outC
	}

}