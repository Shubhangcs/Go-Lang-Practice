package main

import "fmt"

type Details struct{
	Name string
	Email string
	Password string
}

func main(){
	var kiran Details = Details{Name: "Kiran" , Email: "kiran@gmail.com" , Password: "kiran112"}
	fmt.Println(kiran.Email)
}
