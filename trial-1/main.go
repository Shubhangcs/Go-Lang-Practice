package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/demo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(405)
			return
		}
		w.Write([]byte(`<!DOCTYPE html>
		<html>
			<head>
				<title>i am back</title>
			</head>
			<body style="display: flex; justify-content: center; align-items: center; height:100vh;">
				<button onclick="print()" style="width:5rem; height:3rem; background-color: red; color: white; border-radius: 2px;">Print</button>
			</body>
		</html>`))
	})
	log.Println("Server Is Running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
