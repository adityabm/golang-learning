package main

import (
	"firstweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Handler)
	mux.HandleFunc("/post", handler.PostHandler)
	mux.HandleFunc("/siswa", handler.SiswaHandler)

	log.Println("Starting web on port 8822")

	err := http.ListenAndServe(":8822", mux)

	log.Fatal(err)
}