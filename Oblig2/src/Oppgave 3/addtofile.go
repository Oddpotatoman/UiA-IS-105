package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"os/signal"

	"UiA-IS-105/Oblig2/src/Oppgave 3/addfile"
)

func main() {
	d := make(chan os.Signal, 2)
	signal.Notify(d, os.Interrupt,)
	go func() {
		<-d
		fmt.Println("End of Process")
		os.Exit(1)
	}()

	input1 := os.Args[1]
	input2 := os.Args[2]
	input := input1 + " " + input2

	numbers(input)
}

func numbers(input string) {
	done := make(chan bool, 1)

	inputB := []byte(input)

	ioutil.WriteFile("tempfile.txt.lock", inputB, 0777)

	go addfile.Adder(done)

	<-done

	b, err := ioutil.ReadFile("tempfile.txt.lock")
	//"The system cannot find the file specified" dette er en feilmelding som vises dersom filen den går gjennom ikke finnes
	if err != nil {
		fmt.Print(err)
	}

	resultat := string(b)

	fmt.Println(resultat)
}