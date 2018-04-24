package main

import (
	"io/ioutil"
	"fmt"
	"log"
	"strings"
	"strconv"
	"os/signal"
	"os"
)

func main() {
	//Implementer håndtering av SIGINT, her skriver programmet et avslutningsmelding når SIGINT blir motatt
	d := make(chan os.Signal, 1)
	signal.Notify(d, os.Interrupt)

	go func() {
		<-d
		fmt.Println("Nødstopp!!")
		os.Exit(1)
	}()


	addtofile()
	sumfromfile()
	readResult("result.txt")
}
func addtofile() {
	// scanner input og lager fil, deretter sriver den ut filen
	var tallet1 int
	var tallet2 int

	fmt.Println("Skriv inn tall: ")
	fmt.Scan(&tallet1)
	fmt.Println("Skriv inn tall: ")
	fmt.Scan(&tallet2)

	file, err := os.Create("3b.txt")
	if err != nil {
		log.Fatal("Kan ikke opprette fil", err)
	}
	defer file.Close()

	f, err := os.OpenFile("result.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := fmt.Fprintf(f, "%d\n%d", tallet1, tallet2); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func readResult(path string) {
	// leser filen, lagrer siste tall i resultat og konverterer til integer (int)

	data, err := ioutil.ReadFile(path)
	checkErr(err)

	tempData := string(data)
	stringData := strings.Split(tempData, "\n")
	temp := stringData[len(stringData)-2]

	resultat, err := strconv.Atoi(temp)
	checkErr(err)

	fmt.Println("Resultat: ", resultat)
}

