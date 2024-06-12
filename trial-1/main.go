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
			<body>
				<a href="https://github.com/Shubhangcs">Click</a>
			</body>
		</html>`))
	})
	log.Println("Server Is Running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
