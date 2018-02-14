package main

//importerer fmt slik at jeg kan bruke det til å print ut ting
import (
	"fmt"
)

//lager en string hvor det står " € ÷ ¾ dollar " i hexdecimal.
const ascii = "\x22\x20\x80\x20\xF7\x20\xBE\x20\x64\x6F\x6C\x6C\x61\x72\x20\x22"

//lager en main funksjon som kjører når jeg skriver "go run ascii.go" i cmd/powershell/bash osv
//i main kaller den på 2 funksjoner.

func main() {
	iterateOverExtendedASCIIStringLiteral()
	ExtendedASCIIText(ascii)
}

//OPPGAVE 4 a)
//denne funksjonen går igjennom ei løkke som skriver ut extended ascii. Har satt at i er 128 når løkka bagynner.
//Løkka stopper om den er større enn 255. 128-255 representerer extended ascii tegnene.
//Jeg spesifiserer hvilket "format" som skal bli skrevet ut med å skrive: %X, %c og %b. Disse er forskjellige 
//%X betyr base 16 med store bokstaver. Base 16 er hexdecimal.
//%c er unicode.
//%b er base 2 som er binært.
//Fra 0x80 til 0x9f så får vi opp noe som ser ut som firkanter med noe inni når vi bruker powershell, men når vi bruker git bash så får vi ikke opp noe i de feltene

func iterateOverExtendedASCIIStringLiteral(){
	fmt.Println("Oppgave 4a")
	fmt.Println("")
	for i := 128; i <= 255; i++ {
		fmt.Printf("%X %c %b \n", i, i, i)
	}
}

//OPPGAVE 4 b)
//Her bruker jeg const ascii som som har hexdecimalene til det som skal printes ut
//jeg får opp alle tegnene bortsett fra € tegnet. litt usikker på hvorfor, men resten av tegnene er "før" extended ascii
//jeg prøvde å kjøre samme kode på Mac også men fikk samme problem.

func ExtendedASCIIText(tekst string){
		fmt.Println("")
		fmt.Println("Oppgave 4b")
		fmt.Println("")
	j := tekst
	for i := 0; i < len(j); i++ {
		fmt.Printf("%c", j[i])

	}
}
