package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	//dDette er en feilmelding som kommer dersom filen den skal lese ikke finnes.
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
func sumfromfile()  {
	lines, err := readLines("result.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}


	input1 := lines[0]
	input2 := lines[1]

	tallet1,_ := strconv.Atoi(input1)
	tallet2,_ := strconv.Atoi(input2)

	result := tallet1 + tallet2

	file, err := os.OpenFile("result.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Kan ikke åpne fil", err)
	}
	if _, err := fmt.Fprintf(file,"\n%d\n", result); err != nil {
		log.Fatal("Kan ikke skrive fil", err)
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
