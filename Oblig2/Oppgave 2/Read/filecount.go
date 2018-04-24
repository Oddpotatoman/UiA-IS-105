package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//Åpner filen
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	// ioutil.ReadAll() Leser hele filen og printer ut data om bytes lest.
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as hex: %x\n", data)
	fmt.Printf("Data as string: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))
}