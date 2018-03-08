package main

import (
	"os"
	"log"
	"fmt"
	)



func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	fi, err := os.Lstat(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	mode := fi.Mode()
	fmt.Println("Filename")
	fmt.Println(fi.Name())
	if mode.IsRegular() {
		fmt.Println("regular file")
	}else {
		fmt.Println("Is not regular file")
	}
	if mode.IsDir() {
		fmt.Println("directory")
	}else {
		fmt.Println("Is not directory")
	}
	if mode&os.ModeSymlink != 0 {
		fmt.Println("symbolic link")
	}else {
		fmt.Println("Is not Symbbolic link")
	}
	if mode&os.ModeNamedPipe != 0{
		fmt.Println("named pipe")
	}else {
		fmt.Println("Is not named pipe")
	}
	fmt.Println(mode.Perm())
	fmt.Println(fi.Size())

}