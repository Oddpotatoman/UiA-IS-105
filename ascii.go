package main

//importerer fmt slik at jeg kan bruke det til å print ut ting
import (
	"fmt"
)

//lager en main funksjon som kjører når jeg skriver "go run ascii.go" i cmd/powershell/bash osv
//i main kaller den på 2 funksjoner, den første er oppgave a og den andre er oppgave b

func main() {
	iterateOverExtendedASCIIStringLiteral()
}

//OPPGAVE 4 a)
//denne funksjonen går igjennom ei løkke som skriver ut extended ascii. Har satt at i er 128 når løkka bagynner.
//Løkka stopper om den er større enn 255. 128-255 representerer extended ascii tegnene.
//jeg spesifiserer hvilket "format" som skal bli skrevet ut med %X, %c og %b.
//%X betyr base 16 med store bokstaver. Base 16 er hexdecimal.
//%c er unicode.
//%b er base 2 som er binært.

func iterateOverExtendedASCIIStringLiteral(){
	for i := 128; i <= 255; i++ {
		fmt.Printf("%X %c %b \n", i, i, i)
	}
}
