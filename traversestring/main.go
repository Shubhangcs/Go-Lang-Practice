package main

import (
	"io"
	"os"
)

func main() {
    printData()
}

func printData(){
	io.WriteString(os.Stdout , string([]byte("Hello")))
}