package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	  log.Printf("Server is running in port 8000")
	  http.HandleFunc("/login" , func (w http.ResponseWriter , r *http.Request){
		json.NewEncoder(w).Encode("Hello")
	  })
	  log.Fatal(http.ListenAndServe(":8000" , nil))
}

