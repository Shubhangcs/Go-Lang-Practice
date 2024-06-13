package main

import (
	"io"
	"os"
)

func main() {
	Aprintln("hello" + "dhsadjk" , "djksda")
}

func Aprintln(a ...string){
	for i := 0 ; i < len(a) ; i++{
		io.WriteString(os.Stdout , a[i]);
	}
	io.WriteString(os.Stdout , "\n")
}
