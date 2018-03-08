package addfile

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
	"log"
)

func Adder(done chan bool) {
	b, err := ioutil.ReadFile("tempfile.txt.lock")
	//"The system cannot find the file specified" her får vi igen samme feilmeldig som sier at systemet ikke han finne filen
	if err != nil {
		fmt.Print(err)
	}

	inputS := string(b)

	tallS := strings.Fields(inputS)

	tall1S := string(tallS[0])
	tall2S := string(tallS[1])

	tall1, err := strconv.Atoi(tall1S)
	//
	//Feilmeldingen "Invalid syntax" kommer her fordi tall1S ikke er et tall
	if err != nil {
		log.Fatal(err)
	}

	tall2, err := strconv.Atoi(tall2S)
	//Feilmeldingen "Invalid syntax" kommer opp her fordi tall2S ikke er et tall
	if err != nil {
		log.Fatal(err)
	}

	sum := tall1 + tall2

	sumS := strconv.Itoa(sum)

	sumB := []byte(sumS)

	ioutil.WriteFile("tempfile.txt.lock", sumB, 0777)

	done <- true
}