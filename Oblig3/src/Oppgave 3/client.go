package main

import (
	"net"
	"fmt"
	"bufio"
)

//Her tester vi funksjonaliteten til serveren, vi kan enten skrive udp eller tcp for å koble til serveren.
func main() {
	p :=  make([]byte, 2048)
	conn, err := net.Dial("udp", "127.0.0.1:17")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return

	}
	fmt.Fprintf(conn, "What's ut server?")
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}
