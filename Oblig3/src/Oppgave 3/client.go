package main

import (
	"fmt"
	"bufio"
	"os"
	"net"
	"strings"
)

//Her tester vi funksjonaliteten til serveren, vi kan enten skrive udp eller tcp for å koble til serveren.
func main() {
	fmt.Println("Do you wanna connect to the server through TCP or UDP?")
	network := bufio.NewScanner(os.Stdin)
	for network.Scan() {
		connection := strings.ToLower(network.Text())
		p := make([]byte, 2048)
		conn, err := net.Dial(connection, "127.0.0.1:17")
		if err != nil {
			fmt.Printf("Some error %v", err)
			return
		}
		fmt.Fprintf(conn, "Hey server, what's up?")
		_, err = bufio.NewReader(conn).Read(p)
		if err == nil {
			fmt.Printf("%s\n", p)
		} else {
			fmt.Printf("Some error %v\n", err)
		}
		conn.Close()
	}
}

