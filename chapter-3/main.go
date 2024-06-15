package main

// in this tutorial we will learn about declaration of variable in go

import "fmt"

func main(){
	//normal declaration
	var a int8 = 10

	//multiple varible declaration
	var (
		b int16 = 20
		c int32 = 30
	)

	//short hand declaration
	d := 50

	fmt.Printf("a: %T %v\nb: %T %v\nc: %T %v\nd: %T %v\n",a,a,b,b,c,c,d,d)
	
}