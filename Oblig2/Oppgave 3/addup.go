package main

import (

	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	c := make(chan int)
	go readInput(c)
	time.Sleep(5 * 1e9)
	go addUp(c)
	time.Sleep(5 * 1e9)

}

func readInput(c chan int) {

	var tallet1 int
	var tallet2 int
	fmt.Println("Her kan du sette inn ett tall: ")
	fmt.Scan(&tallet1)
	fmt.Println("Her kan du sette inn ett annet tall: ")
	fmt.Scan(&tallet2)

	c <- tallet1 //sender data via channel
	c <- tallet2

	sum := <-c // mottar resultat ifra channel

	fmt.Println("Resultat: ", sum)

}

func addUp(c chan int) {

	tallet1, tallet2 := <-c, <-c // mottar data fra readInput()
	sum := tallet1 + tallet2
	c <- sum // sender resultat tilbake til readInput()

	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan, syscall.SIGINT,)

	exit_chan := make(chan int)

	go func() {
		for {
			s := <-signal_chan
			switch s { // Ctrl+c
			case syscall.SIGINT:
				fmt.Println("Nødstopp!!")

			}
		}
	}()

	code := <-exit_chan
	os.Exit(code)

}