package main

import (
	"fmt"
	"net/http"
	"log"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	log.Println("Client connected: " + r.RemoteAddr)
	defer log.Println("Client disconnected: " + r.RemoteAddr)

	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		log.Println("Couldn`t write ", err.Error())
		return
	}
}

func main() {
	fmt.Println("Starting...")
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
